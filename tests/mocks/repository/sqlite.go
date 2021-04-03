package repository

import (
	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/entity"
	"github.com/stretchr/testify/mock"
)

type SqliteMock struct {
	mock.Mock
}

func (ref *SqliteMock) CreateAccount(account *dto.CreateAccount) (int64, error) {
	args := ref.Called(account)
	return args.Get(0).(int64), args.Error(1)
}
func (ref *SqliteMock) FindAccount(ID int64) (entity.Account, error) {
	args := ref.Called(ID)
	return args.Get(0).(entity.Account), args.Error(1)
}
func (ref *SqliteMock) CreateTransaction(transaction *dto.CreateTransaction) (int64, error) {
	args := ref.Called(transaction)
	return args.Get(0).(int64), args.Error(1)
}
