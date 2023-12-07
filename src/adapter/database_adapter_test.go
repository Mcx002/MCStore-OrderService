package adapter

import (
	"context"
	"mcx002/orderService/src"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestConnect(t *testing.T) {
	// prepare context
	ctx := context.Background()

	t.Run("Should db connected", func(t *testing.T) {
		// load .env file
		err := godotenv.Load("../../.env.test")
		if err != nil {
			t.Errorf("cannot load env :%v", err)
		}

		// get configuration
		appConfig, err := src.NewAppConfig()
		if err != nil {
			t.Errorf("cannot get configuration :%v", err)
		}

		// prepare adapter
		var dbAdapter DatabaseInterface[mongo.Client] = &MongoDbAdapter{}
		err = dbAdapter.Connect(ctx, appConfig.DbUri)
		assert.Nil(t, err)

		// get db client
		client := dbAdapter.GetClient()
		err = client.Ping(ctx, nil)
		assert.Nil(t, err)

		// close the connection
		err = dbAdapter.Close(ctx)
		assert.Nil(t, err)
	})

	t.Run("Should error when trying to connect", func(t *testing.T) {
		// prepare adapter
		var dbAdapter DatabaseInterface[mongo.Client] = &MongoDbAdapter{}
		err := dbAdapter.Connect(ctx, "")
		assert.NotNil(t, err)
	})
}
