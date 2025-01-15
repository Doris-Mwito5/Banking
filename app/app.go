package app

import (
	"log"
	"net/http"
)

func Start() {

	//creae a multiplexer
	mux := http.NewServeMux()
	//define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getCustomers)

	//starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
