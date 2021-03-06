package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Fuerback/transactions-go/domain"
	"github.com/Fuerback/transactions-go/errors"
	"github.com/Fuerback/transactions-go/service"
	"gopkg.in/go-playground/validator.v9"
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
	resp.Header().Set("Content-type", "application/json")
	transactionDTO := new(domain.CreateTransactionDTO)
	err := json.NewDecoder(r.Body).Decode(transactionDTO)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error unmarshalling the request"}`))
		return
	}
	v := validator.New()
	err = v.Struct(transactionDTO)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	transaction, err := transactionService.Create(transactionDTO)
	if err != nil {
		if err == errors.ErrInvalidOperation {
			resp.WriteHeader(http.StatusBadRequest)
		} else {
			resp.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	result, err := json.Marshal(transaction)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error unmarshalling the transaction"}`))
		return
	}
	resp.WriteHeader(http.StatusCreated)
	resp.Write(result)
}
