package repository_test

import (
	"testing"

	"github.com/Fuerback/transactions-go/domain"
	"github.com/Fuerback/transactions-go/repository"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var (
	repo           repository.Repository
	validAccountID int64
)

func init() {
	repo = repository.NewSqlite("../db/transaction_test.db?_foreign_keys=on")
	cleanDB()
}

func TestCreateValidAccount(t *testing.T) {
	repo = repository.NewSqlite("../db/transaction_test.db?_foreign_keys=on")

	account := new(domain.CreateAccountDTO)
	account.DocumentNumber = "5457647"

	ID, err := repo.CreateAccount(account)

	assert.NoError(t, err)
	assert.Greater(t, ID, int64(1))
	validAccountID = ID
}

func TestFindValidAccount(t *testing.T) {
	repo = repository.NewSqlite("../db/transaction_test.db?_foreign_keys=on")

	account, err := repo.FindAccount(validAccountID)

	assert.NoError(t, err)
	assert.NotNil(t, account)
}

func TestCreateValidTransaction(t *testing.T) {
	repo = repository.NewSqlite("../db/transaction_test.db?_foreign_keys=on")

	transaction := new(domain.CreateTransactionDTO)
	transaction.AccountID = validAccountID
	transaction.Amount = 123.78
	transaction.OperationTypeID = 1

	ID, err := repo.CreateTransaction(transaction)

	assert.NoError(t, err)
	assert.Greater(t, ID, int64(1))
}

func TestClearDB(t *testing.T) {
	repo = repository.NewSqlite("../db/transaction_test.db?_foreign_keys=on")

	err := repo.ClearUp()

	assert.NoError(t, err)
}

func cleanDB() {
	repo.ClearUp()
}
