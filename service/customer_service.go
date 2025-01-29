package service

import "github/Doris-Mwito5/banking/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerByID(ID int64) (*domain.Customer, error)
}

//implement  the primary port(business logic)
type customerService struct {
	//insert the dependency
	repo domain.CustomerRepository
}

//helper function to instantiate the service
func NewCustomerService(repository domain.CustomerRepository) customerService {
	return customerService{repo: repository}
}

func (s customerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAllCustomers()
}

func (s customerService) GetCustomerByID(ID int64) (*domain.Customer, error) {
	return s.repo.GetCustomerByID(ID)
}