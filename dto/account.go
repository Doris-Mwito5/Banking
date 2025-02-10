package dto

import "time"

type Account struct {
	ID          int64     `json:"id"`
	CustomerID  string    `json:"customer_id"`
	Pin         string    `json:"pin"`
	AccountType string    `json:"account_type"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
	Amount      float64   `json:"amount"`
}

type AccountRequest struct {
	CustomerID  string    `json:"customer_id"`
	Pin         string    `json:"pin"`
	AccountType string    `json:"account_type"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
	Amount      float64   `json:"amount"`
}
