package server

import (
	"context"
	"fmt"
	"mcx002/orderService/src/proto_gen"
)

func (s *Server) CreateOrder(ctx context.Context, order *proto_gen.OrderDto) (*proto_gen.OrderDto, error) {
	defer ctx.Done()

	// execute create order service
	result, err := s.Service.CreateOrder(order)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	return result, nil
}
