package repository

import (
	"github.com/Fuerback/transactions-go/domain"
	"github.com/stretchr/testify/mock"
)

type SqliteMock struct {
	mock.Mock
}

func (ref *SqliteMock) CreateAccount(account *domain.CreateAccountDTO) (int64, error) {
	args := ref.Called(account)
	return args.Get(0).(int64), args.Error(1)
}
func (ref *SqliteMock) FindAccount(ID int64) (domain.Account, error) {
	args := ref.Called(ID)
	return args.Get(0).(domain.Account), args.Error(1)
}
func (ref *SqliteMock) CreateTransaction(transaction *domain.CreateTransactionDTO) (int64, error) {
	args := ref.Called(transaction)
	return args.Get(0).(int64), args.Error(1)
}
func (ref *SqliteMock) ClearUp() error {
	args := ref.Called()
	return args.Error(0)
}
