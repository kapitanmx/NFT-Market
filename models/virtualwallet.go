package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VirtualWallet struct {
	ID               primitive.ObjectID `bson:"_id"`
	Balance          *float32 `json:"balance"`
	LastTransactions []models.Transaction `json:"last_transactions"`
}
