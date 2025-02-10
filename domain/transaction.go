package domain

import (
	"github/Doris-Mwito5/banking/dto"
	"github/Doris-Mwito5/banking/errors"
	"time"
)

const (
	WITHDRAWAL = "withdraw"
	DEPOSIT    = "deposit"
)

type Transaction struct {
	ID              int64     `json:"id"`
	CustomerID      string    `json:"customer_id"`
	AccountID       int64     `json:"account_id"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	CreatedAt       time.Time `json:"created_at"`
}

func (t Transaction) ToDto() dto.Transaction {
	return dto.Transaction{
		ID:              t.ID,
		CustomerID:      t.CustomerID,
		AccountID:       t.AccountID,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		CreatedAt:       t.CreatedAt,
	}
}

type (
	TransactionRepository interface {
		Save(Transaction) (*Transaction, *errors.AppError)
	}
)

func (t Transaction) IsWithdraw() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t Transaction) IsDeposit() bool {
	return t.TransactionType == DEPOSIT
}
