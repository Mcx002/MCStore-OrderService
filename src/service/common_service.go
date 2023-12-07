package service

import (
	"mcx002/orderService/src/proto_gen"
	"time"
)

func (s Service) GetHealth() *proto_gen.Health {
	return &proto_gen.Health{
		Name:     s.Config.Name,
		Version:  s.Config.Version,
		Lifetime: time.Now().Unix() - s.Config.ServerStartTime.Unix(),
	}
}
