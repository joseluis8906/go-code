package store

import (
	"context"

	"github.com/joseluis8906/go-code/src/pkg/cmp"
	"github.com/joseluis8906/go-code/src/pkg/repository"

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
	}

	Repository struct {
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
					Key:   "Name",
					Value: 1,
				},
				{
					Key:   "City",
					Value: 1,
				},
				{
					Key:   "Address",
					Value: 1,
				},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{
				{
					Key:   "Products.Ref",
					Value: 1,
				},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{
				{
					Key:   "Products.Name",
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
		client:     deps.Conn,
		db:         db,
		collection: collection,
	}

	return repo, nil
}

// Matchig returns the assistant for the given system.
func (r *Repository) Matching(ctx context.Context, criteria cmp.Criteria) repository.Result[Store] {
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

	cursor, err := r.client.Database(r.db).Collection(r.collection).Find(ctx, query, opts)
	if err != nil {
		return repository.Error[Store](err)
	}

	if err = cursor.All(ctx, &result); err != nil {
		return repository.Error[Store](err)
	}

	return repository.Data(result)
}
