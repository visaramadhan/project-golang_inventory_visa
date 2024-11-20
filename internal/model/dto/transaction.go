package model

import "time"

type Transaction struct {
	ID              int       `json:"id"`
	ProductID       int       `json:"product_id"`
	TransactionType string    `json:"transaction_type"`
	Amount          float64   `json:"amount"`
	Description     string    `json:"desc"`
	TimeStamp       time.Time `json:"timestamp"`
}
