package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID         primitive.ObjectID `bson:"_id"`
	Date       string `json:"date" validation:"required"`
	Type       string `json:"type" validation:"required"`
	Payer      string `json:"payer" validation:"required"`
	PayerID    string `json:"payer_id" validation:"required"`
	Receiver   string `json:"receiver" validation:"required"`
	ReceiverID string `json:"receiver_id" validation:"required"`
	Sum        *float32 `json:"sum" validation:"required"`
}
