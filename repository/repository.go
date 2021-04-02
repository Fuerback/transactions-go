package repository

import (
	"database/sql"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/entity"
)

type Repository interface {
	CreateAccount(account *dto.CreateAccount) error
	FindAccount(account *dto.FindAccount) (*entity.Account, error)
	CreateTransaction(transaction *dto.CreateTransaction) error
}

func NewSqlite(db *sql.DB) Repository {
	return &sqlite{
		DB: db,
	}
}
