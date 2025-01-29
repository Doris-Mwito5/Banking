package domain

type Customer struct {
	ID int64
	Name string
	DateOfBirth string
	City string
	ZipCode string
	Status string
}

//Create the secondary port

type CustomerRepository interface {
	FindAllCustomers() ([]Customer, error)
	GetCustomerByID(ID int64) (*Customer, error)
}

//define the adapter/stub
type customerRepository struct {
	customers []Customer
}

