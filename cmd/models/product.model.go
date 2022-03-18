package models

import (
	"models"
)

var (
	productName string   = models.Announcement.Title
	imgs        []string = models.Announcement.Imgs
	tags        []string = models.Announcement.Tags
	category    string   = models.Announcement.Category
	price       int64    = models.Announcement.Price
)

type Product struct {
	ProductName string
	Imgs        []string
	Category    string
	Tags        []string
	Price       int64
}

func (p *Product) SetData() {
	p.ProductName = productName
	p.Imgs = imgs
	p.Tags = tags
	p.Category = category
	p.Price = price
}
