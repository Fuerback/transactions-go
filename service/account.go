package service

import (
	"errors"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/repository"
)

type AccountService interface {
	Find(ID int64) (dto.Account, error)
	Create(account *dto.CreateAccount) error
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

func (s *accountService) Find(ID int64) (dto.Account, error) {
	account, err := repo.FindAccount(ID)
	if err != nil {
		return dto.Account{}, err
	}
	accountDTO, _ := accountParser.ParseDomainToMessage(account)
	return accountDTO, nil
}

func (s *accountService) Create(account *dto.CreateAccount) error {
	err := repo.CreateAccount(account)
	if err != nil {
		return errors.New("Error creating new account in database")
	}
	return nil
}
