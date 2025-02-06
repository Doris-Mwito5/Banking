package domain

import (
	"github/Doris-Mwito5/banking/dto"
	"github/Doris-Mwito5/banking/errors"
	"time"
)

type Account struct {
	ID          int64     `json:"id"`
	CustomerID  string     `json:"customer_id"`
	AccountType string    `json:"account_type"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
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
		ID: a.ID,
		CustomerID: a.CustomerID,
		AccountType: a.AccountType,
		CreatedAt: a.CreatedAt,
		Status: a.statusAsText(),
	}
}

type AccountRepository interface {
	GetAllAccounts() ([]Account, *errors.AppError)
	Save(Account) (*Account, *errors.AppError)
}

type accountRepository struct {
	Account []Account
}

