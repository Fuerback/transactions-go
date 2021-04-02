package service

import (
	"errors"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/entity"
	"github.com/Fuerback/transactions-go/repository"
)

type TransactionService interface {
	Create(transaction *dto.CreateTransaction) error
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

func (s *transactionService) Create(transaction *dto.CreateTransaction) error {
	err := repo.CreateTransaction(transaction)
	if err != nil {
		return errors.New("Error creating new transaction in database")
	}
	return nil
}
