package service

import (
	"github/Doris-Mwito5/banking/domain"
	errors "github/Doris-Mwito5/banking/errors"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *errors.AppError)
	GetCustomerByID(ID string) (*domain.Customer, *errors.AppError)
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

func (s customerService) GetAllCustomers() ([]domain.Customer, *errors.AppError) {
	return s.repo.FindAllCustomers()
}

func (s customerService) GetCustomerByID(ID string) (*domain.Customer, *errors.AppError) {
	return s.repo.GetCustomerByID(ID)
}
