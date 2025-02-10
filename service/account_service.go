package service

import (
	"github/Doris-Mwito5/banking/domain"
	"github/Doris-Mwito5/banking/dto"
	errors "github/Doris-Mwito5/banking/errors"
	"time"
)

type AccountService interface {
	GetAllAccounts() ([]dto.Account, *errors.AppError)
	CreateAccount(request dto.AccountRequest) (*dto.Account, *errors.AppError)
}

type accountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *accountService {
	return &accountService{repo: repository}
}

func (s *accountService) GetAllAccounts() ([]dto.Account, *errors.AppError) {
	a, err := s.repo.GetAllAccounts()
	if err != nil {
		return nil, err
	}
	var response []dto.Account
	for _, acc := range a {
		response = append(response, acc.ToDto())
	}
	return response, nil
}

func (s *accountService) CreateAccount(
	request dto.AccountRequest) (*dto.Account, *errors.AppError) {
	account := domain.Account{
		CustomerID:  request.CustomerID,
		AccountType: request.AccountType,
		CreatedAt:   time.Now(),
		Status:      request.Status,
		Amount: request.Amount,
	}
	newAccount, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToDto()
	return &response, nil
}
