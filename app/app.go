package app

import (
	"fmt"
	"github/Doris-Mwito5/banking/domain"
	"github/Doris-Mwito5/banking/service"
	"os"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//creae a multiplexer
	// mux := http.NewServeMux()
	mux := mux.NewRouter()
	//wiring
	// ch := CustomerHandler{service: service.NewCustomerService(domain.NewcustomerRepository())}
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewcustomerRepoDb())}
	ah := AccountHandler{service: service.NewAccountService(domain.NewaccountRepoDb())}

	//define routes
	mux.HandleFunc("/customers", ch.getAllCustomers)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerByID)
	mux.HandleFunc("/accounts", ah.GetAllAccounts)
	mux.HandleFunc("/account", ah.CreateAccount)

	//starting the server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), mux))
}
