package app

import (
	"encoding/json"
	"github/Doris-Mwito5/banking/dto"
	"github/Doris-Mwito5/banking/service"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, _ := ah.service.GetAllAccounts()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func (ah *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"] 
	var request dto.AccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	request.CustomerID = customerID
	account, appError := ah.service.CreateAccount(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.Message)
		return
	}
	writeResponse(w, http.StatusCreated, account)
}
