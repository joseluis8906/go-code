package store

import (
	"context"
	"fmt"
	"log"

	"github.com/joseluis8906/go-code/src/pkg/repository"
	"go.opentelemetry.io/otel"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.uber.org/fx"
)

// CatalogRepository represents the repository for the catalog.
type (
	Deps struct {
		fx.In

		Conn *mongo.Client
		Logs *log.Logger
	}

	Repository struct {
		logs *log.Logger

		client     *mongo.Client
		db         string
		collection string
	}
)

// NewCatalogRepository creates a new assistant repository instance.
func NewRepository(deps Deps) (*Repository, error) {
	db := "delivery"
	collection := "stores"
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				{
					Key:   Fields().Name,
					Value: 1,
				},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{
				{
					Key:   Fields().Products.Ref,
					Value: 1,
				},
			},
		},
		{
			Keys: bson.D{
				{
					Key:   Fields().Products.Ref,
					Value: 1,
				},
			},
		},
	}

	_, err := deps.Conn.Database(db).
		Collection(collection).
		Indexes().
		CreateMany(context.TODO(), indexes)

	if err != nil {
		return nil, err
	}

	repo := &Repository{
		logs:       deps.Logs,
		client:     deps.Conn,
		db:         db,
		collection: collection,
	}

	return repo, nil
}

// Matchig returns the assistant for the given system.
func (r *Repository) Get(ctx context.Context, criteria repository.Criteria) repository.Result[Store] {
	ctx, span := otel.Tracer("").Start(ctx, "store.Repository.Get")
	defer span.End()

	var result []Store

	opts := options.Find().
		SetLimit(repository.Limit).
		SetSkip(repository.Page(criteria.Page()))

	query := bson.D{
		{
			Key: criteria.Field(),
			Value: bson.D{
				{
					Key:   criteria.Operator(),
					Value: criteria.Value(),
				},
			},
		},
	}

	span.AddEvent(fmt.Sprintf("db.stores.find({\"%s\": {\"$regex\": \"%s\"}})", criteria.Field(), criteria.Value()))
	cursor, err := r.client.Database(r.db).Collection(r.collection).Find(ctx, query, opts)
	if err != nil {
		return repository.Error[Store](fmt.Errorf("searching in mongo: %w", err))
	}

	if err = cursor.All(ctx, &result); err != nil {
		return repository.Error[Store](fmt.Errorf("decoding from mongo cursor: %w", err))
	}

	return repository.Data(result)
}

func (r *Repository) Add(ctx context.Context, aStore Store) error {
	filter := bson.D{
		{
			Key:   Fields().Name,
			Value: aStore.Name.Value,
		},
	}

	_, err := r.client.Database(r.db).
		Collection(r.collection).
		ReplaceOne(ctx, filter, aStore, options.Replace().SetUpsert(true))

	if err != nil {
		return fmt.Errorf("creating or replacing in mongo: %w", err)
	}

	return nil
}
