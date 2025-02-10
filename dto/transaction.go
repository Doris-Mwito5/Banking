package dto

import "time"

const (
	WITHDRAWAL = "withdrawal"
	DEPOSIT  = "deposit"
)

type Transaction struct {
	ID              int64     `json:"id"`
	CustomerID      string    `json:"customer_id"`
	AccountID       int64    `json:"account_id"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	CreatedAt       time.Time `json:"created_at"`
}

type TransactionRequest struct {
	CustomerID      string    `json:"customer_id"`
	AccountID       int64    `json:"account_id"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	CreatedAt       time.Time `json:"created_at"`
}

func (req TransactionRequest) IsWithdrawal() bool {
	return req.TransactionType == WITHDRAWAL
}

func (req TransactionRequest) IsDeposit() bool {
	return req.TransactionType == DEPOSIT
}
