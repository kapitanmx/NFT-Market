package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID primitive.ObjectID `bson:"_id"`
	OrderNo string `json:"order_no" validation:"required"`
	Name    string `json:"name" validation:"required"`
	Vendor string `json:"vendor" validation:"required"`
	Customer string `json:"customer" validation:"required"`
}
