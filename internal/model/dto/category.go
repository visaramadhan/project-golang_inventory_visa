package model

type Category struct {
	ID           int    `json:"id"`
	ProductID    string `json:"product_id"`
	CategoryName string `json:"category_name"`
}
