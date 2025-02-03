package domain

import (
	"github/Doris-Mwito5/banking/dto"
	"github/Doris-Mwito5/banking/errors"
)

type Customer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	Status      string `json:"status"`
}

func(c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func(c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID: c.ID,
		Name: c.Name,
		DateOfBirth: c.DateOfBirth,
		City: c.City,
		ZipCode: c.ZipCode,
		Status: c.statusAsText(),
	}
}

//Create the secondary port

type CustomerRepository interface {
	FindAllCustomers() ([]Customer, *errors.AppError)
	GetCustomerByID(ID string) (*Customer, *errors.AppError)
}

// define the adapter/stub
type customerRepository struct {
	customers []Customer
}
