package service

import (
	"context"
	"mcx002/orderService/src"
	"mcx002/orderService/src/proto_gen"
	"mcx002/orderService/src/repository"

	"github.com/rabbitmq/amqp091-go"
)

type ServiceInterface interface {
	SetContext(context.Context)

	GetHealth()
	CreateOrder(*proto_gen.OrderDto) (*proto_gen.OrderDto, error)
	UpdateCancelOrderEvent(d amqp091.Delivery) error
}

type Service struct {
	Config *src.AppConfig
	ctx    context.Context
	Repo   repository.RepositoryInterface
}

func (s *Service) SetContext(ctx context.Context) {
	s.ctx = ctx
	s.Repo.SetContext(ctx)
}
