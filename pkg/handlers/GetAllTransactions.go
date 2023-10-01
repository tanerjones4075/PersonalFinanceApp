package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tannerjones4075/PersonalFinanceApp/pkg/mocks"
)

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Transactions)
}
