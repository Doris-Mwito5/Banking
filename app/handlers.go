package app

import (
	"github/Doris-Mwito5/banking/service"
	"encoding/json"
	"net/http"
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
	// customers := []Customer{
	// 	{Name: "Doris", City: "Nairobi", ZipCode: "23098"},
	// 	{Name: "Alex", City: "Austria", ZipCode: "45890"},
	// }

	customers, _ := ch.service.GetAllCustomers()
	//setting the content type to json
	w.Header().Add("Content-Type", "application/json")
	//setting the content type to xml
	// w.Header().Add("Content-Type", "application/xml")
	//encoding the customers data to json
	json.NewEncoder(w).Encode(customers)
	//encoding the customers data to xml
	// xml.NewEncoder(w).Encode(customers)

}
