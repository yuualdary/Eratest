package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID                 string `gorm:"primary_key"`
	ProductName        string
	ProductPrice       int
	ProductDescription string
	ProductQuantity    int
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

func (b *Product) TableName() string {
	return "products"
}
