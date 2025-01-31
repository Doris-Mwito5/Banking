package domain

import "github/Doris-Mwito5/banking/errors"

type Customer struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City        string `db:"city"`
	ZipCode     string `db:"zip_code"`
	Status      string `db:"status"`
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
