package app

import (
	"encoding/json"
	"github/Doris-Mwito5/banking/service"
	"net/http"

	"github.com/gorilla/mux"
)

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
	ID := vars["customer_id"]

	customer, err := ch.service.GetCustomerByID(ID)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(code)
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			panic(err)
		}
}
