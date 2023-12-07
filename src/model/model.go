package model

import "go.mongodb.org/mongo-driver/mongo"

type CollectionList struct {
	OrderColl *mongo.Collection
}

func NewCollectionList(client *mongo.Database) *CollectionList {
	return &CollectionList{
		OrderColl: client.Collection("order"),
	}
}
