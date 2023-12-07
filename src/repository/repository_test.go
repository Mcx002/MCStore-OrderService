package repository

import (
	"context"
	"mcx002/orderService/src"
	"mcx002/orderService/src/adapter"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestSetContext(t *testing.T) {
	repo := &Repository{}

	ctx := context.Background()
	repo.SetContext(ctx)

	assert.Equal(t, ctx, repo.ctx)
}

func RepoTestInitiation(t *testing.T) (context.Context, *mongo.Client, *mongo.Database, *src.AppConfig) {
	ctx := context.Background()

	err := godotenv.Load("../../.env.test")
	if err != nil {
		t.Errorf("cannot load env :%v", err)
	}

	appConfig, err := src.NewAppConfig()
	if err != nil {
		t.Errorf("failed to load configuration :%v", err)
	}

	dbAdapter := &adapter.MongoDbAdapter{}
	err = dbAdapter.Connect(ctx, appConfig.DbUri)
	if err != nil {
		t.Errorf("failed to connect to db :%v", err)
	}

	client := dbAdapter.GetClient()
	db := client.Database(appConfig.DbName)

	return ctx, client, db, appConfig
}
