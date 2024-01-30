package nosql

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"go.uber.org/fx"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Params struct {
	fx.In

	Config *viper.Viper
	Logger *log.Logger
}

func New(params Params) *mongo.Client {
	opts := options.Client().
		ApplyURI(params.Config.GetString("mongodb.uri"))

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		params.Logger.Fatalf("connecting mongo: %v", err)
	}

	return client
}
