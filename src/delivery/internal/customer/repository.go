package customer

import (
	"context"

	"github.com/joseluis8906/go-code/src/pkg/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.uber.org/fx"
)

// CustomerRepository represents an in-memory repository for customers.
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

// NewCustomerRepository returns a new instance of CustomerInMemoryRepository.
func NewRepository(deps Deps) (*Repository, error) {
	db := "delivery"
	collection := "customers"
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				{
					Key:   "Email",
					Value: 1,
				},
			},
			Options: options.Index().SetUnique(true),
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

// Matching returns the customer for the given email.
func (r *Repository) Get(ctx context.Context, criteria repository.Criteria) repository.Result[Customer] {
	var result []Customer

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
		return repository.Error[Customer](err)
	}

	if err = cursor.All(ctx, &result); err != nil {
		return repository.Error[Customer](err)
	}

	return repository.Data(result)
}

func (r *Repository) Add(ctx context.Context, aCustomer Customer) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{
		{
			Key:   EmailField,
			Value: aCustomer.Email.Value,
		},
	}

	_, err := r.client.Database(r.db).
		Collection(r.collection).
		UpdateOne(ctx, filter, aCustomer, opts)

	return err
}
