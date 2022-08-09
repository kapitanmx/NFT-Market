package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ProductName primitive.ObjectID `bson:"_id` 
	Imgs        []string `json:"imgs"`
	Category    string `json:"category`
	Tags        []string `json:"tags"`
	Price       *float32 `json:"price"`
}

