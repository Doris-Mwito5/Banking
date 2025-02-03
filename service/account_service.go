package service

import (
	"github/Doris-Mwito5/banking/domain"
	"github/Doris-Mwito5/banking/dto"
	errors "github/Doris-Mwito5/banking/errors"
)

type AccountService interface {
	GetAllAccounts() ([]dto.Account, *errors.AppError)
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
