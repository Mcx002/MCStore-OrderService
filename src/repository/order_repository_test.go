package repository

import (
	"fmt"
	"mcx002/orderService/src/adapter"
	"mcx002/orderService/src/model"
	"mcx002/orderService/src/proto_gen"
	"testing"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestGetOrderCount(t *testing.T) {
	ctx, client, db, appConfig := RepoTestInitiation(t)

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should failed to get count documents", func(mt *mtest.T) {
		r := &Repository{
			Collections: &model.CollectionList{
				OrderColl: mt.Coll,
			},
			Client: client,
			Config: appConfig,
		}
		r.SetContext(ctx)

		mt.AddMockResponses(bson.D{{"ok", 0}})

		result, err := r.GetOrderCountToday()

		assert.NotNil(t, err)
		assert.Equal(t, result, int64(0))
	})

	t.Run("Should success get count document", func(t *testing.T) {
		collection := model.NewCollectionList(db)
		r := &Repository{
			Collections: collection,
			Client:      client,
			Config:      appConfig,
		}
		r.SetContext(ctx)

		orders := []interface{}{
			&model.Order{
				Id:        primitive.NewObjectID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			&model.Order{
				Id:        primitive.NewObjectID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			&model.Order{
				Id:        primitive.NewObjectID(),
				CreatedAt: time.Date(2023, 04, 14, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2023, 04, 14, 0, 0, 0, 0, time.UTC),
			},
		}
		_, err := collection.OrderColl.InsertMany(ctx, orders)
		if err != nil {
			t.Errorf("failed to insert many")
		}

		result, err := r.GetOrderCountToday()

		assert.Nil(t, err)
		assert.Equal(t, result, int64(2))

		err = collection.OrderColl.Drop(ctx)
		if err != nil {
			t.Errorf("failed to drop order coll: %v", err)
		}
	})
}

func TestInsertOrder(t *testing.T) {
	ctx, client, db, appConfig := RepoTestInitiation(t)

	r := &Repository{
		Collections: model.NewCollectionList(db),
		Client:      client,
		Config:      appConfig,
	}
	r.SetContext(ctx)

	t.Run("Should failed to publish the food log message", func(t *testing.T) {
		mba := &adapter.MessageBrokerAdapter{
			AppConfig: appConfig,
			Ch:        &MessageBrokerMockChannel{},
			Conn:      &MessageBrokerMockConnection{},
		}

		r.SetBroker(mba)

		order := &model.Order{
			Id: primitive.NewObjectID(),
		}
		defer r.Collections.OrderColl.DeleteOne(ctx, bson.D{{"_id", order.Id}})

		// prepare mock publish
		MessageBrokerMockChannelExchangeDeclare = func(name, kind string, durable, autoDelete, internal, noWait bool, args amqp091.Table) error {
			return fmt.Errorf("error")
		}

		_, err := r.InsertOrder(order, []model.OrderItem{
			{
				Qty:  1,
				Note: "test",
				Variety: map[string]string{
					"color": "yellow",
				},
				FoodSnapshot: &proto_gen.FoodDto{
					Id: "foodId",
				},
			},
		})

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "failed to publish the food log message")
	})

	t.Run("Should success insert order", func(t *testing.T) {
		mba := &adapter.MessageBrokerAdapter{
			AppConfig: appConfig,
		}
		mba.Connect()

		r.SetBroker(mba)

		order := &model.Order{
			Id: primitive.NewObjectID(),
		}
		defer r.Collections.OrderColl.DeleteOne(ctx, bson.D{{"_id", order.Id}})

		id, err := r.InsertOrder(order, []model.OrderItem{
			{
				Qty:  1,
				Note: "test",
				Variety: map[string]string{
					"color": "yellow",
				},
				FoodSnapshot: &proto_gen.FoodDto{
					Id: "foodId",
				},
			},
		})

		assert.Nil(t, err)
		assert.NotEqual(t, id, "")
	})
}
