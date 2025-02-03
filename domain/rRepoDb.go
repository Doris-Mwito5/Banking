package domain

import (
	"database/sql"
	"log"

	"github/Doris-Mwito5/banking/errors"
	"github/Doris-Mwito5/banking/logger"

	_ "github.com/lib/pq"
)

type customerRepoDb struct {
	db *sql.DB
}

func (d customerRepoDb) FindAllCustomers() ([]Customer, *errors.AppError) {

	//sql query to get customers
	FindAllCustomersSQL := `SELECT id, name, date_of_birth, city, zip_code, status FROM customers`
	//query the sql via the db client

	rows, err := d.db.Query(FindAllCustomersSQL)
	if err != nil {
		logger.Error("error querying customers")
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	defer rows.Close()
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.DateOfBirth,
			&c.City,
			&c.ZipCode,
			&c.Status,
		)
		if err != nil {
			logger.Error("scan row err: %v")
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}
		customers = append(customers, c)
	}
	//if no err return the list of customers
	return customers, nil
}

func (d customerRepoDb) GetCustomerByID(ID string) (*Customer, *errors.AppError) {
	getCustomerByIdSQL := `SELECT id, name, date_of_birth, city, zip_code, status FROM customers WHERE id = $1`

	row := d.db.QueryRow(
		getCustomerByIdSQL,
		ID,
	)
	var c Customer
	err := row.Scan(
		&c.ID,
		&c.Name,
		&c.DateOfBirth,
		&c.City,
		&c.ZipCode,
		&c.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("Customer Not found")
		} else {
			logger.Error("scanning row error: %v")
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}

	}
	return &c, nil
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
