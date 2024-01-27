package product

import (
	"context"

	"github.com/joseluis8906/go-code/src/pkg/cmp"
	"github.com/joseluis8906/go-code/src/pkg/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CatalogRepository represents the repository for the catalog.
type Repository struct {
	client     *mongo.Client
	db         string
	collection string
}

// NewCatalogRepository creates a new assistant repository instance.
func NewRepository(ctx context.Context, uri string) (*Repository, error) {
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	db := "delivery"
	collection := "products"
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{"Ref", 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{{"Name", 1}},
		},
	}

	_, err = conn.Database(db).
		Collection(collection).
		Indexes().
		CreateMany(ctx, indexes)

	if err != nil {
		return nil, err
	}

	repo := &Repository{
		client:     conn,
		db:         db,
		collection: collection,
	}

	return repo, nil
}

// Matchig returns the assistant for the given system.
func (r *Repository) Matching(ctx context.Context, criteria cmp.Criteria) repository.Result[Product] {
	var result []Product

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
		return repository.Error[Product](err)
	}

	if err = cursor.All(ctx, &result); err != nil {
		return repository.Error[Product](err)
	}

	return repository.Data(result)
}
