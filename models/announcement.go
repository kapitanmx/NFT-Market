package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Announcement struct {
	ID             primitive.ObjectID `bson:"_id"`
	Date           string `json:"date" validate: "required"`
	ExpDate        string `json:"exp_date" validate: "required"`
	Title          *string `json:"title" validate: "required"`
	Desc           string `json:"desc" validate: "required, min=10, max=1000"`
	Imgs           []string `json:"imgs"`
	Price          *float64 `json:"price" validate: "required"`
	AdvertiserName string `json:"advertiser_name" validate: "required"`
	AdvertiserID   string `json:"advertiser_id" validate: "required"`
	Category       string `json:"category" validate: "required"`
	Tags           []string `json:"tags"`
}
