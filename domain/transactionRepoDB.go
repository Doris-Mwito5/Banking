package domain

import (
	"database/sql"
	"fmt"
	"github/Doris-Mwito5/banking/errors"
	"github/Doris-Mwito5/banking/logger"
	"strings"
)

type transactionRepoDB struct {
	db            *sql.DB
	accountRepoDb *accountRepoDb
}

func NewTransactionRepoDB(db *sql.DB, accountRepoDb *accountRepoDb) *transactionRepoDB {
	return &transactionRepoDB{db: db, accountRepoDb: accountRepoDb}
}

func (d *transactionRepoDB) Save(t Transaction) (*Transaction, *errors.AppError) {
	//Begin a transaction
	tx, err := d.db.Begin()
	if err != nil {
		logger.Error("error while starting a new bank transaction: %+v")
		return nil, errors.NewUnexpectedError("unexpected database error")
	}

	// Insert bank account transaction
	var transactionID int64
	err = tx.QueryRow(
		"INSERT INTO transactions (customer_id, account_id, amount, transaction_type, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id",
		t.CustomerID, t.AccountID, t.Amount, t.TransactionType, t.CreatedAt,
	).Scan(&transactionID)

	if err != nil {
		tx.Rollback()
		logger.Error("error while inserting transaction: %+v")
		return nil, errors.NewUnexpectedError("unexpected database error")
	}

	// Updating account balance
	// if t.IsWithdraw() {
	// 	_, err = tx.Exec("UPDATE accounts SET amount = amount - $1 WHERE id = $2", t.Amount, t.AccountID)
	// } else {
	// 	_, err = tx.Exec("UPDATE accounts SET amount = amount + $1 WHERE id = $2", t.Amount, t.AccountID)
	// }

	logger.Info(fmt.Sprintf("Transaction Type: %s, Amount: %.2f, Account ID: %d", t.TransactionType, t.Amount, t.AccountID))

	if strings.ToLower(t.TransactionType) == "withdrawal" {
		logger.Info("Processing withdrawal...")
		_, err = tx.Exec("UPDATE accounts SET amount = amount - $1 WHERE id = $2", t.Amount, t.AccountID)
	} else {
		logger.Info("Processing deposit...")
		_, err = tx.Exec("UPDATE accounts SET amount = amount + $1 WHERE id = $2", t.Amount, t.AccountID)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("error while updating account balance: %+v")
		return nil, errors.NewUnexpectedError("unexpected database error")
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		logger.Error("error while committing transaction: %+v")
		return nil, errors.NewUnexpectedError("unexpected database error")
	}

	// Retrieve latest account info
	account, AppError := d.accountRepoDb.GetAccountByID(t.AccountID)
	if AppError != nil {
		return nil, AppError
	}

	// Update the transaction struct with the latest balance
	t.ID = transactionID
	t.Amount = account.Amount

	return &t, nil
}
