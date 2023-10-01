package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tannerjones4075/PersonalFinanceApp/pkg/mocks"
)

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, transaction := range mocks.Transactions {
		if transaction.Id == id {
			mocks.Transactions = append(mocks.Transactions[:index], mocks.Transactions[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Transaction Deleted")
			break
		}
	}
}
