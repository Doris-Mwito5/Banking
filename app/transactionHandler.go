package app

import (
	"encoding/json"
	"github/Doris-Mwito5/banking/dto"
	"github/Doris-Mwito5/banking/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (th *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountIDStr := vars["account_id"]
	customerID := vars["customer_id"]
	accountID, err := strconv.ParseInt(accountIDStr, 10, 64)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid Account ID format")
		return
	}
	var request dto.TransactionRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	request.AccountID = accountID
	request.CustomerID = customerID
	transaction, appError := th.service.CreateTransaction(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.Message)
		return
	}
	writeResponse(w, http.StatusCreated, transaction)
}
