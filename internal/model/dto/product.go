package model

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity string `json:"qty"`
	Price    string `json:"price"`
}
