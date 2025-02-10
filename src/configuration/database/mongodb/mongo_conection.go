package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	MONGODB_URL = "MONGODB_URL"
)

func NewMongoConnection() (*mongo.Database, error) {
	mongdb_uri := os.Getenv(MONGODB_URL)
	ctx := context.Background()
	client, err := mongo.Connect(options.Client().ApplyURI(mongdb_uri))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			return nil, err
		}
	}()
	return err
}
