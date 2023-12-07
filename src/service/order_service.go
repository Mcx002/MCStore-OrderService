package service

import (
	"fmt"
	"mcx002/orderService/src/model"
	"mcx002/orderService/src/proto_gen"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Service) UpdateCancelOrderEvent(d amqp.Delivery) error {
	fmt.Println("Hello baraya")
	return nil
}

func (s Service) CreateOrder(order *proto_gen.OrderDto) (*proto_gen.OrderDto, error) {
	// prepare order items
	items := []model.OrderItem{}
	for _, item := range order.Items {
		items = append(items, model.OrderItem{
			Qty:          item.Qty,
			Variety:      item.Variety,
			Note:         item.Note,
			FoodSnapshot: item.FoodSnapshot,
		})
	}

	// get order count today
	count, err := s.Repo.GetOrderCountToday()
	if err != nil {
		return nil, fmt.Errorf("failed to get order count: %v", err)
	}

	// prepare invoice id
	now := time.Now()
	invoiceId := fmt.Sprintf(
		"INV-%04d%02d%02d-%02d%02d%02d-%07d",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(),
		count+1)

	// prepare order model
	model := &model.Order{
		Id:             primitive.NewObjectID(),
		UserSnapshot:   order.User,
		SellerSnapshot: order.Seller,
		Invoice: &model.Invoice{
			Id:             invoiceId,
			ActualCost:     order.Invoice.ActualCost,
			FinalCost:      order.Invoice.FinalCost,
			AdditionalCost: order.Invoice.AdditionalCost,
			Status:         proto_gen.InvoiceDto_PENDING,
		},
		Items:     items,
		Version:   1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// persist insert order
	_, err = s.Repo.InsertOrder(model, items)
	if err != nil {
		return nil, fmt.Errorf("failed to insert order: %v", err)
	}

	return composeOrderDto(model), nil
}

func composeOrderDto(model *model.Order) *proto_gen.OrderDto {
	// prepare item
	items := []*proto_gen.OrderDto_Item{}
	for _, item := range model.Items {
		items = append(items, &proto_gen.OrderDto_Item{
			Qty:          item.Qty,
			Variety:      item.Variety,
			Note:         item.Note,
			FoodSnapshot: item.FoodSnapshot,
		})
	}

	return &proto_gen.OrderDto{
		Id:     model.Id.Hex(),
		User:   model.UserSnapshot,
		Seller: model.SellerSnapshot,
		Items:  items,
		Invoice: &proto_gen.InvoiceDto{
			Id:             model.Invoice.Id,
			ActualCost:     model.Invoice.ActualCost,
			FinalCost:      model.Invoice.FinalCost,
			AdditionalCost: model.Invoice.AdditionalCost,
			Status:         model.Invoice.Status,
		},
		Version:   model.Version,
		CreatedAt: model.CreatedAt.Unix(),
		UpdatedAt: model.UpdatedAt.Unix(),
	}
}
