package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/errors"
	"github.com/Fuerback/transactions-go/service"
)

var transactionService service.TransactionService

type TransactionController interface {
	CreateTransaction(resp http.ResponseWriter, r *http.Request)
}

type transactionController struct{}

func NewTransactionController(service service.TransactionService) TransactionController {
	transactionService = service
	return &transactionController{}
}

func (c *transactionController) CreateTransaction(resp http.ResponseWriter, r *http.Request) {
	transactionDTO := new(dto.CreateTransaction)
	err := json.NewDecoder(r.Body).Decode(transactionDTO)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error unmarshalling the request"}`))
		return
	}
	err = transactionService.Create(transactionDTO)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	resp.WriteHeader(http.StatusCreated)
}
