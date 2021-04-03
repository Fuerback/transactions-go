package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/service"
	"github.com/Fuerback/transactions-go/tests/mocks/repository"
	"github.com/stretchr/testify/suite"
)

const id int64 = 1

type transactionSuite struct {
	suite.Suite
	ctx     context.Context
	service service.TransactionService
	repo    *repository.SqliteMock
}

func TestTransactionServer(t *testing.T) {
	suite.Run(t, &transactionSuite{
		ctx: context.Background(),
	})
}

func (ref *transactionSuite) SetupTest() {
	ref.repo = new(repository.SqliteMock)
	ref.service = service.NewTransactionService(ref.repo)
}

func (ref *transactionSuite) TestCreateTransaction_Success() {
	t := getCreateTransactionDTO(2)

	ref.repo.On("CreateTransaction", t).Return(id, nil).Once()

	transaction, err := ref.service.Create(t)
	ref.NoError(err)
	ref.NotNil(transaction)
	ref.Equal(transaction.ID, id)
	ref.Equal(transaction.AccountID, id)
	ref.Equal(transaction.AccountID, id)
}

func (ref *transactionSuite) TestCreateTransactionNegativeAmount_Success() {

	i := 1
	for i <= 3 {
		t := getCreateTransactionDTO(i)

		ref.repo.On("CreateTransaction", t).Return(id, nil).Once()

		transaction, err := ref.service.Create(t)
		ref.NoError(err)
		ref.NotNil(transaction)
		ref.Equal(transaction.ID, id)
		ref.Negative(transaction.Amount)
		i++
	}

}

func (ref *transactionSuite) TestCreateTransactionPositiveAmount_Success() {

	t := getCreateTransactionDTO(4)

	ref.repo.On("CreateTransaction", t).Return(id, nil).Once()

	transaction, err := ref.service.Create(t)
	ref.NoError(err)
	ref.NotNil(transaction)
	ref.Equal(transaction.ID, id)
	ref.Positive(transaction.Amount)

}

func (ref *transactionSuite) TestCreateTransactionInvalidOperation() {

	t := getCreateTransactionDTO(5)

	ref.repo.On("CreateTransaction", t).Return(id, nil).Once()

	_, err := ref.service.Create(t)
	ref.Error(err)
}

func (ref *transactionSuite) TestCreateTransactionErrorOnPersist() {

	t := getCreateTransactionDTO(1)
	fakeErr := errors.New("some error")

	ref.repo.On("CreateTransaction", t).Return(id, fakeErr).Once()

	_, err := ref.service.Create(t)
	ref.Error(err)
}

func getCreateTransactionDTO(operation int) *dto.CreateTransaction {
	return &dto.CreateTransaction{
		AccountID:       id,
		OperationTypeID: operation,
		Amount:          120.56,
	}
}
