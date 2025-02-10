package domain

import (
	"database/sql"
	"github/Doris-Mwito5/banking/errors"
	"github/Doris-Mwito5/banking/logger"
)

type accountRepoDb struct {
	db *sql.DB
}

func NewaccountRepoDb(db *sql.DB) *accountRepoDb {
	return &accountRepoDb{db: db}
}

func (d *accountRepoDb) GetAllAccounts() ([]Account, *errors.AppError) {
	getAllAccountsSQL := `SELECT id, customer_id, account_type, pin, created_at, status, amount FROM accounts`
	rows, err := d.db.Query(getAllAccountsSQL)
	if err != nil {
		logger.Error("error querying accounts")
		return nil, errors.NewUnexpectedError("unexpected error occurred")
	}
	defer rows.Close()
	accounts := make([]Account, 0)
	for rows.Next() {
		var account Account
		err := rows.Scan(
			&account.ID,
			&account.CustomerID,
			&account.Pin,
			&account.AccountType,
			&account.CreatedAt,
			&account.Status,
			&account.Amount,
		)
		if err != nil {
			logger.Error("scan row err")
			return nil, errors.NewUnexpectedError("unexpected error occurred")
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (d *accountRepoDb) Save(a Account) (*Account, *errors.AppError) {
	createAccountSQL := `INSERT INTO accounts (customer_id, pin, account_type, created_at, status, amount) 
                         VALUES($1, $2, $3, $4, $5, $6) RETURNING id`

	err := d.db.QueryRow(
		createAccountSQL,
		a.CustomerID,
		a.Pin,
		a.AccountType,
		a.CreatedAt,
		a.Status,
		a.Amount,
	).Scan(&a.ID)

	if err != nil {
		logger.Error("error creating account: %+v")
		return nil, errors.NewUnexpectedError("failed to create account: " + err.Error())
	}

	return &a, nil
}

func (d *accountRepoDb) GetAccountByID(ID int64) (*Account, *errors.AppError) {
	getAccountByIDSQL := `SELECT customer_id, pin, account_type, created_at, amount, status FROM accounts WHERE ID = $1`

	row := d.db.QueryRow(
		getAccountByIDSQL,
		ID,
	)
	var a Account
	err := row.Scan(
		&a.CustomerID,
		&a.Pin,
		&a.AccountType,
		&a.CreatedAt,
		&a.Amount,
		&a.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("account not found")
		} else {
			logger.Error("scanning row error: %v")
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}
	}
	return &a, nil
}
