package models

import (
	"models"

	"github.com/google/uuid"
)

type UserCart struct {
	ID         string
	Products   []models.Product
	TotalPrice int
}

func (uc *UserCart) SetID() error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	uc.ID = id.String()
	return nil
}

func (uc *UserCart) AddProduct(p *Product) {
	append(uc.Products, p)
}

func (uc *UserCart) SetTotalPrice() {
	productsArr := uc.Products
	sum := 0
	for i := 0; i < len(productsArr); i++ {
		sum += productsArr[i].Price
	}
	uc.TotalPrice = sum
}
