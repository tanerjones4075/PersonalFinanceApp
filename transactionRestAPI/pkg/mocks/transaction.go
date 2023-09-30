package mocks

import (
	"github.com/tannerjones4075/PersonalFinanceApp/transactionRestApi/pkg/models"

) 
var Articles = []models.Transaction{
	{Id: "1", TransDate: "01/1/2023", Merchant: "Amazon", Amount: 100, AccountNumber 8817},
	{Id: "2", TransDate: "02/1/2023", Merchant: "Taco Bell", Amount: 100, AccountNumber 8817},
}

// https://dev.to/janirefdez/create-a-rest-api-with-go-1j52