package service

import (
	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/entity"
)

type (
	AccountParser     struct{}
	TransactionParser struct{}
)

func (ref AccountParser) ParseDomainToMessage(e entity.Account) (dto.Account, error) {
	account := dto.Account{
		ID:             e.ID,
		DocumentNumber: e.DocumentNumber,
	}
	return account, nil
}

func (ref TransactionParser) ParseDomainToMessage(ID int64, t *dto.CreateTransaction) (dto.Transaction, error) {
	transactionDto := dto.Transaction{
		ID:              ID,
		AccountID:       t.AccountID,
		OperationTypeID: t.OperationTypeID,
		Amount:          t.Amount,
	}
	return transactionDto, nil
}
