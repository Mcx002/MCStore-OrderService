package server

import (
	"context"
	"fmt"
	"mcx002/orderService/src/adapter"
	"mcx002/orderService/src/proto_gen"

	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) ListenToMessageBroker() (*adapter.MessageBrokerAdapter, error) {
	mbAdapter := &adapter.MessageBrokerAdapter{
		AppConfig: s.Config,
	}
	err := mbAdapter.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to message broker")
	}

	// Add listener
	mbAdapter.ListenTo("update-cancel-order", adapter.ListenOption{}, s.Service.UpdateCancelOrderEvent)

	return mbAdapter, nil
}

func (s *Server) GetHealth(ctx context.Context, in *empty.Empty) (*proto_gen.Health, error) {
	defer ctx.Done()

	health := s.Service.GetHealth()

	return health, nil
}
