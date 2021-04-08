package service_test

import (
	"errors"
	"testing"

	"github.com/Fuerback/transactions-go/domain"
	"github.com/Fuerback/transactions-go/service"
	"github.com/Fuerback/transactions-go/tests/mocks/repository"
	"github.com/stretchr/testify/assert"
)

const id int64 = 1

var (
	transactionService service.TransactionService
	repo               *repository.SqliteMock
)

func init() {
	repo = new(repository.SqliteMock)
	transactionService = service.NewTransactionService(repo)
}

func TestCreateTransaction_Success(t *testing.T) {
	dto := getCreateTransactionDTO(2)

	repo.On("CreateTransaction", dto).Return(id, nil).Once()

	transaction, err := transactionService.Create(dto)

	repo.AssertExpectations(t)
	assert.NotNil(t, transaction)
	assert.NoError(t, err)
	assert.Equal(t, id, transaction.ID)
	assert.Equal(t, id, transaction.AccountID)
}

func TestCreateTransactionNegativeAmount_Success(t *testing.T) {

	i := 1
	for i <= 3 {
		dto := getCreateTransactionDTO(i)

		repo.On("CreateTransaction", dto).Return(id, nil).Once()

		transaction, err := transactionService.Create(dto)

		repo.AssertExpectations(t)
		assert.NoError(t, err)
		assert.NotNil(t, transaction)
		assert.Equal(t, id, transaction.ID)
		assert.Negative(t, transaction.Amount)
		i++
	}

}

func TestCreateTransactionPositiveAmount_Success(t *testing.T) {

	dto := getCreateTransactionDTO(4)

	repo.On("CreateTransaction", dto).Return(id, nil).Once()

	transaction, err := transactionService.Create(dto)

	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, id, transaction.ID)
	assert.Positive(t, transaction.Amount)

}

func TestCreateTransactionInvalidOperation(t *testing.T) {

	dto := getCreateTransactionDTO(5)

	_, err := transactionService.Create(dto)

	assert.Error(t, err)
}

func TestCreateTransactionErrorOnPersist(t *testing.T) {

	dto := getCreateTransactionDTO(1)
	fakeErr := errors.New("some error")

	repo.On("CreateTransaction", dto).Return(id, fakeErr).Once()

	_, err := transactionService.Create(dto)

	repo.AssertExpectations(t)
	assert.Error(t, err)
}

func getCreateTransactionDTO(operation int) *domain.CreateTransactionDTO {
	return &domain.CreateTransactionDTO{
		AccountID:       id,
		OperationTypeID: operation,
		Amount:          120.56,
	}
}
