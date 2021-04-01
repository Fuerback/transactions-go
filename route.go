package main

import (
	"encoding/json"
	"net/http"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/entity"
)

var (
	createAccountParser     CreateAccountParser
	createTransactionParser CreateTransactionParser
)

var accounts []entity.Account
var transactions []entity.Transaction

func init() {
	accounts = []entity.Account{}
	transactions = []entity.Transaction{}
}

func createAccount(resp http.ResponseWriter, r *http.Request) {
	var accountDTO dto.CreateAccount
	err := json.NewDecoder(r.Body).Decode(&accountDTO)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error unmarshalling the request"}`))
		return
	}
	entity, err := createAccountParser.ParseMessageToDomain(accountDTO)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error parsing the request"}`))
		return
	}
	accounts = append(accounts, entity)
	resp.WriteHeader(http.StatusOK)
}

func getAccount(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var accountDTO dto.GetAccount
	err := json.NewDecoder(r.Body).Decode(&accountDTO)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error unmarshalling the request"}`))
		return
	}
	result, err := json.Marshal(accounts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error unmarshalling the accounts"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func createTransaction(resp http.ResponseWriter, r *http.Request) {
	var transactionDTO dto.CreateTransaction
	err := json.NewDecoder(r.Body).Decode(&transactionDTO)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error unmarshalling the request"}`))
		return
	}
	entity, err := createTransactionParser.ParseMessageToDomain(transactionDTO)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error parsing the request"}`))
		return
	}
	transactions = append(transactions, entity)
	resp.WriteHeader(http.StatusOK)
}
