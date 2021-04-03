package service

import (
	"errors"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/entity"
	"github.com/Fuerback/transactions-go/repository"
)

type TransactionService interface {
	Create(transaction *dto.CreateTransaction) (dto.Transaction, error)
}

type transactionService struct{}

var (
	transactionParser TransactionParser
	repo1             repository.Repository
	transactions      []entity.Transaction
)

func init() {
	transactions = []entity.Transaction{}
}

func NewTransactionService(r repository.Repository) TransactionService {
	repo = r
	return &transactionService{}
}

func (s *transactionService) Create(t *dto.CreateTransaction) (dto.Transaction, error) {
	ID, err := repo.CreateTransaction(t)
	if err != nil {
		return dto.Transaction{}, errors.New("Error creating new transaction in database")
	}
	transactionDto := dto.Transaction{ID: ID, AccountID: t.AccountID, OperationTypeID: t.OperationTypeID, Amount: t.Amount}
	return transactionDto, nil
}
