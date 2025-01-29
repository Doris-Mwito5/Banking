package app

import (
	"github/Doris-Mwito5/banking/domain"
	"github/Doris-Mwito5/banking/service"

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

	//define routes
	mux.HandleFunc("/customers", ch.getAllCustomers)

	//starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
