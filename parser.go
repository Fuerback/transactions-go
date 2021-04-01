package main

import (
	"time"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/entity"
)

type (
	CreateAccountParser     struct{}
	CreateTransactionParser struct{}
)

func (ref CreateAccountParser) ParseMessageToDomain(d dto.CreateAccount) (entity.Account, error) {
	// TODO: o db que deve criar um novo ID
	account := entity.Account{
		ID:             time.Now().Minute(),
		DocumentNumber: d.DocumentNumber,
	}

	return account, nil
}

func (ref CreateTransactionParser) ParseMessageToDomain(d dto.CreateTransaction) (entity.Transaction, error) {
	// TODO: o db que deve criar um novo ID
	transaction := entity.Transaction{
		ID:            time.Now().Minute(),
		AccountID:     d.AccountID,
		OperationType: d.OperationTypeID,
		Amount:        d.Amount,
		EventDate:     time.Now().Format(time.RFC3339),
	}

	return transaction, nil
}
