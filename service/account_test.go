package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/entity"
	e "github.com/Fuerback/transactions-go/errors"
	"github.com/Fuerback/transactions-go/service"
	"github.com/Fuerback/transactions-go/tests/mocks/repository"
	"github.com/stretchr/testify/suite"
)

const documentNumber = "32451348003"

type accountSuite struct {
	suite.Suite
	ctx     context.Context
	service service.AccountService
	repo    *repository.SqliteMock
}

func TestAccountServer(t *testing.T) {
	suite.Run(t, &accountSuite{
		ctx: context.Background(),
	})
}

func (ref *accountSuite) SetupTest() {
	ref.repo = new(repository.SqliteMock)
	ref.service = service.NewAccountService(ref.repo)
}

func (ref *accountSuite) TestCreateAccount_Success() {
	a := getCreateAccountDTO()

	ref.repo.On("CreateAccount", a).Return(id, nil).Once()

	account, err := ref.service.Create(a)
	ref.NoError(err)
	ref.NotNil(account)
	ref.Equal(account.ID, id)
	ref.Equal(account.DocumentNumber, documentNumber)
}

func (ref *accountSuite) TestFindAccount_Success() {

	ref.repo.On("FindAccount", id).Return(getAccount(), nil).Once()

	account, err := ref.service.Find(id)
	ref.NoError(err)
	ref.NotNil(account)
	ref.Equal(account.ID, id)
	ref.Equal(account.DocumentNumber, documentNumber)
}

func (ref *accountSuite) TestDoNotFindAccount() {

	fakeErr := errors.New("some error")
	ref.repo.On("FindAccount", id).Return(entity.Account{}, fakeErr).Once()

	_, err := ref.service.Find(id)
	ref.NotNil(err)
}

func (ref *accountSuite) TestCreateAccountInvalidDocumentNumber() {
	a := &dto.CreateAccount{
		DocumentNumber: "documentNumber03973",
	}

	_, err := ref.service.Create(a)
	ref.NotNil(err)
	ref.Equal(err.Error(), e.ErrInvalidDocumentNumber.Error())
}

func (ref *accountSuite) TestCreateAccountErrorOnPersist() {
	a := getCreateAccountDTO()

	fakeErr := errors.New("some error")
	ref.repo.On("CreateAccount", a).Return(id, fakeErr).Once()

	_, err := ref.service.Create(a)
	ref.NotNil(err)
}

func getCreateAccountDTO() *dto.CreateAccount {
	return &dto.CreateAccount{
		DocumentNumber: documentNumber,
	}
}

func getAccount() entity.Account {
	return entity.Account{
		ID:             id,
		DocumentNumber: documentNumber,
	}
}
