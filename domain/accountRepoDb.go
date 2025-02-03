package domain

import (
	"database/sql"
	"github/Doris-Mwito5/banking/errors"
	"github/Doris-Mwito5/banking/logger"
	"log"
)

type accountRepoDb struct {
	db *sql.DB
}

func NewaccountRepoDb() *accountRepoDb {
	//db connection
	connStr := "user=root dbname=postgres sslmode=disable password=random123 host=localhost port=5434"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//check if the db connection is active
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	return &accountRepoDb{db}
}

func (d *accountRepoDb) GetAllAccounts() ([]Account, *errors.AppError) {
	getAllAccountsSQL := `SELECT id, customer_id, account_type, created_at, status FROM accounts`
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
			&account.AccountType,
			&account.CreatedAt,
			&account.Status,
		)
		if err != nil {
			logger.Error("scan row err")
			return nil, errors.NewUnexpectedError("unexpected error occurred")
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
