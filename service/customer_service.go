package service

import (
	"github/Doris-Mwito5/banking/domain"
	"github/Doris-Mwito5/banking/dto"
	errors "github/Doris-Mwito5/banking/errors"
)

type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errors.AppError)
	GetCustomerByID(ID string) (*dto.CustomerResponse, *errors.AppError)
}

// implement  the primary port(business logic)
type customerService struct {
	//insert the dependency
	repo domain.CustomerRepository
}

// helper function to instantiate the service
func NewCustomerService(repository domain.CustomerRepository) customerService {
	return customerService{repo: repository}
}

func (s customerService) GetAllCustomers() ([]dto.CustomerResponse, *errors.AppError) {
	c, err := s.repo.FindAllCustomers()
	if err != nil {
		return nil, err
	}
	var response []dto.CustomerResponse
	for _, c := range c {
		response = append(response, c.ToDto())
	}
	return response, nil
}

func (s customerService) GetCustomerByID(ID string) (*dto.CustomerResponse, *errors.AppError) {
	c, err := s.repo.GetCustomerByID(ID)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}
