package domain

import (
	"github/Doris-Mwito5/banking/dto"
	"github/Doris-Mwito5/banking/errors"
	"time"
)

type Account struct {
	ID          int64     `json:"id"`
	CustomerID  string    `json:"customer_id"`
	Pin         string    `json:"pin"`
	AccountType string    `json:"account_type"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
	Amount      float64   `json:"amount"`
}

func (a Account) statusAsText() string {
	statusAsText := "active"
	if a.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (a Account) ToDto() dto.Account {
	return dto.Account{
		ID:          a.ID,
		CustomerID:  a.CustomerID,
		Pin:         a.Pin,
		AccountType: a.AccountType,
		CreatedAt:   a.CreatedAt,
		Status:      a.statusAsText(),
		Amount:      a.Amount,
	}
}

type AccountRepository interface {
	GetAllAccounts() ([]Account, *errors.AppError)
	Save(Account) (*Account, *errors.AppError)
	GetAccountByID(ID int64) (*Account, *errors.AppError)
}

type accountRepository struct {
	Account []Account
}

