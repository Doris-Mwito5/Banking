package domain

type Customer struct {
	ID int
	Name string
	DateOfBirth string
	City string
	ZipCode string
	Status string
}

//Create the secondary port

type CustomerRepository interface {
	FindAllCustomers() ([]Customer, error)
}

//define the adapter/stub
type customerRepository struct {
	customers []Customer
}

