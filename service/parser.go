package service

import (
	"github.com/Fuerback/transactions-go/domain"
)

type (
	AccountParser     struct{}
	TransactionParser struct{}
)

func (ref AccountParser) ParseAccountEntityToAccountDTO(e domain.Account) (domain.AccountDTO, error) {
	account := domain.AccountDTO{
		ID:             e.ID,
		DocumentNumber: e.DocumentNumber,
	}
	return account, nil
}

func (ref AccountParser) ParseCreateAccountToAccount(ID int64, a *domain.CreateAccountDTO) (domain.AccountDTO, error) {
	account := domain.AccountDTO{
		ID:             ID,
		DocumentNumber: a.DocumentNumber,
	}
	return account, nil
}

func (ref TransactionParser) ParseCreateTransactionToTransaction(ID int64, t *domain.CreateTransactionDTO) (domain.TransactionDTO, error) {
	transactionDto := domain.TransactionDTO{
		ID:              ID,
		AccountID:       t.AccountID,
		OperationTypeID: t.OperationTypeID,
		Amount:          t.Amount,
	}
	return transactionDto, nil
}
