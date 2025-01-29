package app

import (
	"encoding/json"
	"fmt"
	"github/Doris-Mwito5/banking/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" city:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

// define the hndler which has a dependency of the service
type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()
	//setting the content type to json
	w.Header().Add("Content-Type", "application/json")
	//encoding the customers data to json
	json.NewEncoder(w).Encode(customers)

}

func (ch *CustomerHandler) getCustomerByID(w http.ResponseWriter, r *http.Request) {
	//fetch the customer id
	vars := mux.Vars(r)
	idStr := vars["customer_id"]

	// Convert ID to int64
	ID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}
	customer, err := ch.service.GetCustomerByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
