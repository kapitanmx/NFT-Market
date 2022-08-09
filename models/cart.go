package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserCart struct {
	ID         primitive.ObjectID `bson:"_id"`
	Products   []models.Product `json:"products"`
	TotalPrice *float32 `json:"total_price"`
}

