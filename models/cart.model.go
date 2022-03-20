package models

import (
	"models"
)

type UserCart struct {
	CartID     string
	Products   []models.Product
	TotalPrice int
}

func (uc *UserCart) AddProduct(product ...Product) {
	len := 100
	productList := make([]Product, len)
	append(productList, product...)
	append(uc.Products, productList)
}

func (uc *UserCart) SetSum() {
	productsArr := uc.Products
	sum := 0
	for i := 0; i < len(productsArr); i++ {
		sum += productsArr[i].Price
	}
	uc.TotalPrice = sum
}
