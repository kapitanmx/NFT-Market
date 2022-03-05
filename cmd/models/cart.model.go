package models

import (
	"models"
)

type UserCart struct {
	Products   []models.Product
	TotalPrice int
}

func (uc *UserCart) SetSum() {
	productsArr := uc.Products
	sum := 0
	for i := 0; i < len(productsArr); i++ {
		sum += productsArr[i].Price
	}
	uc.TotalPrice = sum
}
