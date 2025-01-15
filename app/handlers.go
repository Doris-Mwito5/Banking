package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" city:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Doris", City: "Nairobi", ZipCode: "23098"},
		{Name: "Alex", City: "Austria", ZipCode: "45890"},
	}
	//setting the content type to json
	w.Header().Add("Content-Type", "application/json")
	//setting the content type to xml
	// w.Header().Add("Content-Type", "application/xml")
	//encoding the customers data to json
	json.NewEncoder(w).Encode(customers)
	//encoding the customers data to xml
	xml.NewEncoder(w).Encode(customers)

}
