package service

import (
	"errors"
	"strconv"

	"github.com/Fuerback/transactions-go/domain"
	e "github.com/Fuerback/transactions-go/errors"
	"github.com/Fuerback/transactions-go/repository"
)

type AccountService interface {
	Find(ID int64) (domain.AccountDTO, error)
	Create(account *domain.CreateAccountDTO) (domain.AccountDTO, error)
}

type accountService struct{}

var (
	accountParser AccountParser
	repo          repository.Repository
)

func NewAccountService(r repository.Repository) AccountService {
	repo = r
	return &accountService{}
}

func (s *accountService) Find(ID int64) (domain.AccountDTO, error) {
	account, err := repo.FindAccount(ID)
	if err != nil {
		return domain.AccountDTO{}, err
	}
	accountDTO, _ := accountParser.ParseAccountEntityToAccountDTO(account)
	return accountDTO, nil
}

func (s *accountService) Create(account *domain.CreateAccountDTO) (domain.AccountDTO, error) {
	err := s.validateDocumentNumber(account)
	if err != nil {
		return domain.AccountDTO{}, err
	}
	ID, err := repo.CreateAccount(account)
	if err != nil {
		return domain.AccountDTO{}, errors.New("Error creating new account in database")
	}
	accountDto, _ := accountParser.ParseCreateAccountToAccount(ID, account)
	return accountDto, nil
}

func (s *accountService) validateDocumentNumber(account *domain.CreateAccountDTO) error {
	_, err := strconv.Atoi(account.DocumentNumber)
	if err != nil {
		return e.ErrInvalidDocumentNumber
	}
	return nil
}
