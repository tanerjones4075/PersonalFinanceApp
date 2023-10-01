package mocks

import (
	"github.com/tannerjones4075/PersonalFinanceApp/pkg/models"
)

var Transactions = []models.Transaction{
	{Id: "1", TransDate: "01/1/2023", Merchant: "Amazon", Amount: 100, AccountNumber: 8817, Category: "Shopping"},
	{Id: "2", TransDate: "02/1/2023", Merchant: "Taco Bell", Amount: 100, AccountNumber: 8817, Category: "Shopping"},
}

// https://dev.to/janirefdez/create-a-rest-api-with-go-1j52
