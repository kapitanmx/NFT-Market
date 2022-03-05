package models

import (
	"checkers"

	"github.com/google/uuid"
)

type User struct {
	ID           string
	UserName     string
	UserLastName string
	Email        string
	Password     string
	Country      string
	Street       string
	HouseNumber  string
	PostalCode   string
	City         string
	Sex          string
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
