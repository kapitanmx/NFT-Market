package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Base
	UserID      primitive.ObjectID `bson:"_id"`
	Token		*string `json:"token" validate:"required"`
	Name        *string `json:"user_name,omitempty" validate:"required"`
	LastName    *string `json:"user_last_name,omitempty" validate:"required"`
	Avatar		*string `json:"avatar"`
	Email       *string `json:"email,omitempty" validate:"required"`
	Password    *string `json:"password,omitempty" validate:"required"`
	Country     *string `json:"country,omitempty" validate:"required"`
	Street      *string `json:"street,omitempty" validate:"required"`
	HouseNumber *string `json:"house_number,omitempty" validate:"required"`
	PostalCode  *string `json:"postal_code,omitempty" validate:"required"`
	City        *string `json:"city,omitempty" validate:"required"`
	Sex         *string `json:"sex,omitempty" validate:"required"`
}

