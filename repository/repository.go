package repository

import "github.com/Fuerback/transactions-go/domain"

type Repository interface {
	CreateAccount(account *domain.CreateAccountDTO) (int64, error)
	FindAccount(ID int64) (domain.Account, error)
	CreateTransaction(transaction *domain.CreateTransactionDTO) (int64, error)
	ClearUp() error
}

func NewSqlite(dbFilePath string) Repository {
	return &sqlite{DBFilePath: dbFilePath}
}
