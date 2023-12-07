package server

import (
	"context"
	"mcx002/orderService/src"
	pb "mcx002/orderService/src/proto_gen"
	"mcx002/orderService/src/service"
)

type Server struct {
	pb.UnimplementedOrderServer
	Config  *src.AppConfig
	Service *service.Service
}

func (s *Server) SetContext(ctx context.Context) {
	s.Service.SetContext(ctx)
}
