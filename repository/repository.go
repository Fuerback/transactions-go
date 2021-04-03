package repository

import (
	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/entity"
)

type Repository interface {
	CreateAccount(account *dto.CreateAccount) (int64, error)
	FindAccount(ID int64) (entity.Account, error)
	CreateTransaction(transaction *dto.CreateTransaction) (int64, error)
}

func NewSqlite() Repository {
	return &sqlite{}
}
