package app

import (
	"database/sql"
	"fmt"
	"github/Doris-Mwito5/banking/domain"
	"github/Doris-Mwito5/banking/service"
	"os"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Initialize the database connection
func InitDB() *sql.DB {
	connStr := "user=root dbname=postgres sslmode=disable password=random123 host=localhost port=5434"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// Check if the connection is active
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection is not active: ", err)
	}

	return db
}

func Start() {

	//creae a multiplexer
	// mux := http.NewServeMux()
	mux := mux.NewRouter()
	//wiring

	db := InitDB()

	// Initialize repositories (inject db)
	accountRepo := domain.NewaccountRepoDb(db)                      
	transactionRepo := domain.NewTransactionRepoDB(db, accountRepo) 

	// Initialize services (inject repositories)
	accountService := service.NewAccountService(accountRepo)
	transactionService := service.NewTransactionService(transactionRepo, accountRepo) 

	// ch := CustomerHandler{service: service.NewCustomerService(domain.NewcustomerRepository())}
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewcustomerRepoDb())}
	ah := AccountHandler{service: accountService}
	th := TransactionHandler{service: transactionService}
	//define routes
	mux.HandleFunc("/customers", ch.getAllCustomers)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerByID)
	mux.HandleFunc("/accounts", ah.GetAllAccounts)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.CreateAccount)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", th.CreateTransaction)

	//starting the server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), mux))
}
