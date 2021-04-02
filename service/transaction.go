package service

import (
	"errors"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/entity"
)

type TransactionService interface {
	Create(transaction *dto.CreateTransaction) error
}

type transactionService struct{}

var (
	createTransactionParser CreateTransactionParser
	transactions            []entity.Transaction
)

func init() {
	transactions = []entity.Transaction{}
}

func NewTransactionService() TransactionService {
	return &transactionService{}
}

func (s *transactionService) Create(transaction *dto.CreateTransaction) error {
	entity, err := createTransactionParser.ParseMessageToDomain(transaction)
	if err != nil {
		return errors.New("Error parsing CreateTransactionDTO to transaction entity")
	}
	transactions = append(transactions, entity)
	return nil
}
