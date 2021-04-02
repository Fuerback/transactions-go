package service

import (
	"errors"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/entity"
)

type AccountService interface {
	Find(account *dto.FindAccount) ([]entity.Account, error)
	Create(account *dto.CreateAccount) error
}

type accountService struct{}

var (
	createAccountParser CreateAccountParser
	accounts            []entity.Account
)

func init() {
	accounts = []entity.Account{}
}

func NewAccountService() AccountService {
	return &accountService{}
}

func (s *accountService) Find(account *dto.FindAccount) ([]entity.Account, error) {
	// find account in db
	return accounts, nil
}

func (s *accountService) Create(account *dto.CreateAccount) error {
	entity, err := createAccountParser.ParseMessageToDomain(account)
	if err != nil {
		return errors.New("Error parsing CreateAccountDTO to account entity")
	}
	accounts = append(accounts, entity)
	return nil
}
