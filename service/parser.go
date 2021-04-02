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
