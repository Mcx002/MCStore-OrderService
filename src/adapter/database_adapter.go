package adapter

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseInterface[T any] interface {
	GetClient() *T
	Connect(context.Context, string) error
	Close(context.Context) error
}

type MongoDbAdapter struct {
	client *mongo.Client
}

func (db *MongoDbAdapter) Connect(ctx context.Context, dbUri string) error {
	// create mongodb connection
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		return fmt.Errorf("failed when trying to connect to db: %v", err)
	}

	db.client = client

	return nil
}

func (db MongoDbAdapter) GetClient() *mongo.Client {
	return db.client
}

func (db MongoDbAdapter) Close(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}
