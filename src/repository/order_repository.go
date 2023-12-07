package repository

import (
	"encoding/json"
	"fmt"
	"mcx002/orderService/src/adapter"
	"mcx002/orderService/src/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Repository) GetOrderCountToday() (int64, error) {
	// prepare filter this day
	now := time.Now()
	filter := bson.D{
		{"createdAt", bson.D{
			{"$gt", time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())},
			{"$lt", time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 0, 0, now.Location())},
		}},
	}

	// persist count documents
	result, err := r.Collections.OrderColl.CountDocuments(r.ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("failed to get count documents")
	}

	return result, nil
}

func (r Repository) InsertOrder(order *model.Order, orderItems []model.OrderItem) (string, error) {
	// start transaction
	session, err := r.Client.StartSession()
	if err != nil {
		return "", fmt.Errorf("failed to start a session transaction")
	}
	defer session.EndSession(r.ctx)

	// persist data
	id := ""
	err = mongo.WithSession(r.ctx, session, func(ctx mongo.SessionContext) error {
		// persist order
		result, err := r.Collections.OrderColl.InsertOne(ctx, order)
		if err != nil {
			return fmt.Errorf("failed to insert the order: %v", err)
		}
		id = result.InsertedID.(primitive.ObjectID).Hex()

		// marshal order items
		body, err := json.Marshal(orderItems)
		if err != nil {
			return fmt.Errorf("failed to convert orderItems to byte: %v", err)
		}

		// inform food service
		err = r.broker.Publish(
			r.ctx,
			r.Config.FoodServiceExchange,
			r.Config.FoodServiceRouteAddFoodChangeLog,
			body,
			adapter.PublishOption{
				Durable:    true,
				Persistent: true,
			},
		)
		if err != nil {
			return fmt.Errorf("failed to publish the food log message: %v", err)
		}

		return err
	})

	return id, err
}
