package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tannerjones4075/PersonalFinanceApp/pkg/mocks"
	"github.com/tannerjones4075/PersonalFinanceApp/pkg/models"
)

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedTransaction models.Transaction
	json.Unmarshal(body, &updatedTransaction)

	for index, transaction := range mocks.Transactions {

		if transaction.Id == id {
			transaction.TransDate = updatedTransaction.TransDate
			transaction.Merchant = updatedTransaction.Merchant
			transaction.Amount = updatedTransaction.Amount

			mocks.Transactions[index] = transaction

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
