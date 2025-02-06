package dto

import "time"

type Account struct {
	ID          int64     `json:"id"`
	CustomerID  string     `json:"customer_id"`
	AccountType string    `json:"account_type"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
}

type AccountRequest struct {
	CustomerID  string     `json:"customer_id"`
	AccountType string    `json:"account_type"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
}
