package models

import (
	"checkers"

	"github.com/google/uuid"
)

type User struct {
	ID           string `json:"id,omitempty"`
	UserName     string `json:"username,omitempty" validate:"required"`
	UserLastName string `json:"userlastname,omitempty" validate:"required"`
	Email        string `json:"email,omitempty" validate:"required"`
	Password     string `json:"password,omitempty" validate:"required"`
	Country      string `json:"country,omitempty" validate:"required"`
	Street       string `json:"street,omitempty" validate:"required"`
	HouseNumber  string `json:"housenumber,omitempty" validate:"required"`
	PostalCode   string `json:"postalcode,omitempty" validate:"required"`
	City         string `json:"city,omitempty" validate:"required"`
	Sex          string `json:"sex,omitempty"`
	IsAdult      bool
}

func (u *User) CreateUser(
	userName,
	userLastName,
	userEmail,
	password,
	country,
	street,
	houseNumber,
	postalCode,
	city,
	sex string) (*User, error) {
	id, err := uuid.UUID()
	if err != nil {
		return nil, err
	}
	user := &User{
		ID:           id,
		UserName:     userName,
		UserLastName: userLastName,
		Email:        userEmail,
		Password:     password,
		Country:      country,
		Street:       street,
		HouseNumber:  houseNumber,
		PostalCode:   postalCode,
		City:         city,
		Sex:          sex,
	}
	if checkers.IsUserEmpty() {
		return nil, err
	}
	return user, nil
}

func (u *User) GetUserID() string {
	return u.ID
}
