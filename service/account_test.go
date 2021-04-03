package service_test

import (
	"errors"
	"testing"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/entity"
	e "github.com/Fuerback/transactions-go/errors"
	"github.com/Fuerback/transactions-go/service"
	"github.com/Fuerback/transactions-go/tests/mocks/repository"
	"github.com/stretchr/testify/assert"
)

const documentNumber = "32451348003"

var (
	accountService service.AccountService
	repoAccount    *repository.SqliteMock
)

func init() {
	repo = new(repository.SqliteMock)
	accountService = service.NewAccountService(repo)
}

func TestCreateAccount_Success(t *testing.T) {
	a := getCreateAccountDTO()

	repo.On("CreateAccount", a).Return(id, nil).Once()

	account, err := accountService.Create(a)

	repo.AssertExpectations(t)
	assert.NotNil(t, account)
	assert.NoError(t, err)
	assert.Equal(t, id, account.ID)
	assert.Equal(t, documentNumber, account.DocumentNumber)
}

func TestFindAccount_Success(t *testing.T) {

	repo.On("FindAccount", id).Return(getAccount(), nil).Once()

	account, err := accountService.Find(id)

	repo.AssertExpectations(t)
	assert.NotNil(t, account)
	assert.NoError(t, err)
	assert.Equal(t, id, account.ID)
	assert.Equal(t, documentNumber, account.DocumentNumber)
}

func TestDoNotFindAccount(t *testing.T) {

	fakeErr := errors.New("some error")
	repo.On("FindAccount", id).Return(entity.Account{}, fakeErr).Once()

	_, err := accountService.Find(id)

	repo.AssertExpectations(t)
	assert.NotNil(t, err)
}

func TestCreateAccountInvalidDocumentNumber(t *testing.T) {
	a := &dto.CreateAccount{
		DocumentNumber: "documentNumber03973",
	}

	_, err := accountService.Create(a)

	repo.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Equal(t, e.ErrInvalidDocumentNumber.Error(), err.Error())
}

func TestCreateAccountErrorOnPersist(t *testing.T) {
	a := getCreateAccountDTO()

	fakeErr := errors.New("some error")
	repo.On("CreateAccount", a).Return(id, fakeErr).Once()

	_, err := accountService.Create(a)

	repo.AssertExpectations(t)
	assert.NotNil(t, err)
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
