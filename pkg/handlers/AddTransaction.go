package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tannerjones4075/PersonalFinanceApp/pkg/mocks"
	"github.com/tannerjones4075/PersonalFinanceApp/pkg/models"

	"github.com/google/uuid"
)

func AddTransaction(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var transaction models.Transaction
	json.Unmarshal(body, &transaction)

	transaction.Id = (uuid.New()).String() // Generates a random number for the id
	mocks.Transactions = append(mocks.Transactions, transaction)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Transaction Created")
}
