package service

import (
	"errors"

	"github.com/Fuerback/transactions-go/domain"
	"github.com/Fuerback/transactions-go/repository"
)

type TransactionService interface {
	Create(transaction *domain.CreateTransactionDTO) (domain.TransactionDTO, error)
}

type transactionService struct{}

var (
	transactionParser TransactionParser
	repo1             repository.Repository
)

func NewTransactionService(r repository.Repository) TransactionService {
	repo = r
	return &transactionService{}
}

func (s *transactionService) Create(t *domain.CreateTransactionDTO) (domain.TransactionDTO, error) {
	err := s.transformAmountByOperationID(t)
	if err != nil {
		return domain.TransactionDTO{}, err
	}
	ID, err := repo.CreateTransaction(t)
	if err != nil {
		return domain.TransactionDTO{}, errors.New("Error creating new transaction in database")
	}
	transactionDto, _ := transactionParser.ParseCreateTransactionToTransaction(ID, t)
	return transactionDto, nil
}

func (s *transactionService) transformAmountByOperationID(t *domain.CreateTransactionDTO) error {
	isNegative, err := domain.IsNegative(t.OperationTypeID)
	if err != nil {
		return err
	}
	if isNegative {
		t.Amount = -t.Amount
	}
	return nil
}
