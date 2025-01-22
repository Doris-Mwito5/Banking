package domain

//reate the function to retrieve all the customers
func (s customerRepository) FindAllCustomers() ([]Customer, error) {
	return s.customers, nil
}

//create a new function that creates dummy data/instantiates the functions
func NewcustomerRepository() customerRepository {
	customers := []Customer {
		{ID: 1000, Name: "Doris", DateOfBirth: "1999-11-28", City: "Nairobi", ZipCode: "606000", Status: "1"},
		{ID: 1001, Name: "Ruth", DateOfBirth: "1983-09-11", City: "Meru", ZipCode: "670677", Status: "1"},
	}
	return customerRepository{customers: customers}
}