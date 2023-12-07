package model

import (
	"mcx002/orderService/src/proto_gen"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id             primitive.ObjectID `bson:"_id"`
	UserSnapshot   *proto_gen.UserDto `bson:"userSnapshot"`
	SellerSnapshot *proto_gen.UserDto `bson:"sellerSnapshot"`
	Invoice        *Invoice           `bson:"invoice"`
	Items          []OrderItem        `bson:"items"`
	Version        int32              `bson:"version"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}

type OrderItem struct {
	Qty          int32              `bson:"qty"`
	Variety      map[string]string  `bson:"variety"`
	Note         string             `bson:"note"`
	FoodSnapshot *proto_gen.FoodDto `bson:"foodSnapshot"`
}

type Invoice struct {
	Id             string                      `bson:"id"`
	ActualCost     float64                     `bson:"actualCost"`
	FinalCost      float64                     `bson:"finalCost"`
	AdditionalCost map[string]float64          `bson:"additionalCost"`
	Status         proto_gen.InvoiceDto_Status `bson:"status"`
}
