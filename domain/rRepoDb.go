package domain

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)
type customerRepoDb struct{
	db *sql.DB
}

func (d customerRepoDb) FindAllCustomers() ([]Customer, error) {
	
	//sql query to get customers
	FindAllCustomersSQL := `SELECT id, name, date_of_birth, city, zip_code, status FROM customers`
	//query the sql via the db client

	rows, err := d.db.Query(FindAllCustomersSQL)
	if err != nil {
		log.Printf("error querying customers: %v", err)
		return nil, err
	}
	//if no err, loop through rows

	customers := make([]Customer, 0)
	for rows.Next() {
		//create a var for storing the queried data
		var c Customer
		//retrieve data using scan method
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.DateOfBirth,
			&c.City,
			&c.ZipCode,
			&c.Status,
		)
		if err != nil {
			log.Printf("scan row err: %v", err)
			return nil, err
		}
		defer rows.Close()
		//if no err, we will first customer and we will store them in the named var below
		customers = append(customers, c)
	}
	//if no err return the list of customers
	return customers, nil
}

func NewcustomerRepoDb() customerRepoDb {
	//db connection
	connStr := "user=root dbname=postgres sslmode=disable password=random123 host=localhost port=5434"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// Check if the connection is alive
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	return customerRepoDb{db}
}