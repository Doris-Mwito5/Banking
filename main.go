package main

import (
	"database/sql"
	"github/Doris-Mwito5/banking/app"
	"github/Doris-Mwito5/banking/logger"
	"log"
)

func main() {
	// Initialize DB connection
	connStr := "user=root dbname=postgres sslmode=disable password=random123 host=localhost port=5434"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Check if the DB connection is active
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	logger.Info("Starting the application...")
	app.Start()
}
