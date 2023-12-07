package repository

import (
	"context"
	"mcx002/orderService/src"
	"mcx002/orderService/src/adapter"
	"mcx002/orderService/src/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryInterface interface {
	SetContext(context.Context)
	InsertOrder(*model.Order, []model.OrderItem) (string, error)
	GetOrderCountToday() (int64, error)
}

type Repository struct {
	ctx         context.Context
	Collections *model.CollectionList
	Client      *mongo.Client
	Config      *src.AppConfig
	broker      *adapter.MessageBrokerAdapter
}

func (r *Repository) SetContext(ctx context.Context) {
	r.ctx = ctx
}

func (r *Repository) SetBroker(broker *adapter.MessageBrokerAdapter) {
	r.broker = broker
}
