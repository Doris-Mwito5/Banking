package app

import (
	"encoding/json"
	"github/Doris-Mwito5/banking/service"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, _ := ah.service.GetAllAccounts()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

