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
	err := s.transformAmountByOperationID(t)
	if err != nil {
		return dto.Transaction{}, err
	}
	ID, err := repo.CreateTransaction(t)
	if err != nil {
		return dto.Transaction{}, errors.New("Error creating new transaction in database")
	}
	transactionDto, _ := transactionParser.ParseDomainToMessage(ID, t)
	return transactionDto, nil
}

func (s *transactionService) transformAmountByOperationID(t *dto.CreateTransaction) error {
	isNegative, err := entity.IsNegative(t.OperationTypeID)
	if err != nil {
		return err
	}
	if isNegative {
		t.Amount = -t.Amount
	}
	return nil
}
