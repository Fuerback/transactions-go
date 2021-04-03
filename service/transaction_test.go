package service_test

import (
	"context"
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

func TestHanlderServer(t *testing.T) {
	suite.Run(t, &transactionSuite{
		ctx: context.Background(),
	})
}

func (ref *transactionSuite) SetupTest() {
	ref.repo = new(repository.SqliteMock)
	ref.service = service.NewTransactionService(ref.repo)
}

func (ref *transactionSuite) TestCreateTransaction_Success() {
	t := getCreateTransactionDTO()

	ref.repo.On("CreateTransaction", t).Return(id, nil).Once()

	transaction, err := ref.service.Create(t)
	ref.NoError(err)
	ref.NotNil(transaction)
	ref.Equal(transaction.ID, id)
	ref.Equal(transaction.AccountID, id)
	ref.Equal(transaction.AccountID, id)
}

func getCreateTransactionDTO() *dto.CreateTransaction {
	return &dto.CreateTransaction{
		AccountID:       id,
		OperationTypeID: 2,
		Amount:          120.56,
	}
}
