package service

import (
	"github/Doris-Mwito5/banking/domain"
	"github/Doris-Mwito5/banking/dto"
	errors "github/Doris-Mwito5/banking/errors"
	"time"
)

type TransactionService interface {
	CreateTransaction(req dto.TransactionRequest) (*dto.Transaction, *errors.AppError)
}

type transactionService struct {
	transactionRepo domain.TransactionRepository
	accountRepo     domain.AccountRepository
}

func NewTransactionService(
	tRepo domain.TransactionRepository,
	aRepo domain.AccountRepository) *transactionService {
	return &transactionService{
		transactionRepo: tRepo,
		accountRepo:     aRepo,
	}
}

func (s *transactionService) CreateTransaction(req dto.TransactionRequest) (*dto.Transaction, *errors.AppError) {
	// Validate the request
	err := s.Validate(req)
	if err != nil {
		return nil, err
	}

	if req.IsWithdrawal() {
		account, err := s.accountRepo.GetAccountByID(req.AccountID)
		if err != nil {
			return nil, err
		}
		if account.Amount < req.Amount {
			return nil, errors.NewValidationError("insufficient balance in the account")
		}
	}

	// Create and save the transaction
	transaction := domain.Transaction{
		CustomerID:      req.CustomerID,
		AccountID:       req.AccountID,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		CreatedAt:       time.Now(),
	}

	newTransaction, err := s.transactionRepo.Save(transaction)
	if err != nil {
		return nil, err
	}

	// Return transaction ID and updated balance
	response := newTransaction.ToDto()
	return &response, nil
}

func (s *transactionService) Validate(req dto.TransactionRequest) *errors.AppError {
	// Validate transaction type
	if req.TransactionType != dto.DEPOSIT && req.TransactionType != dto.WITHDRAWAL {
		return errors.NewValidationError("transaction type can only be 'deposit' or 'withdrawal'")
	}

	// Ensure amount is positive
	if req.Amount < 0 {
		return errors.NewValidationError("transaction amount cannot be less than zero")
	}

	return nil
}
