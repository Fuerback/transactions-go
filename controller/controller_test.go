package controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Fuerback/transactions-go/controller"
	"github.com/Fuerback/transactions-go/domain"
	"github.com/Fuerback/transactions-go/repository"
	"github.com/Fuerback/transactions-go/service"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var (
	repo                  repository.Repository
	accountService        service.AccountService
	accountController     controller.AccountController
	transactionService    service.TransactionService
	transactionController controller.TransactionController
	validAccountID        int64
)

func init() {
	repo = repository.NewSqlite("../db/transaction_test.db?_foreign_keys=on")
	cleanDB()
	accountService = service.NewAccountService(repo)
	accountController = controller.NewAccountController(accountService)
	transactionService = service.NewTransactionService(repo)
	transactionController = controller.NewTransactionController(transactionService)
}

func TestFindInvalidAccount(t *testing.T) {
	var testList = []string{"1", "test"}

	for _, test := range testList {
		req, _ := http.NewRequest("GET", "/account", nil)

		handler := http.HandlerFunc(accountController.FindAccount)

		vars := map[string]string{
			"id": test,
		}

		req = mux.SetURLVars(req, vars)

		response := httptest.NewRecorder()

		handler.ServeHTTP(response, req)

		status := response.Code
		if status != http.StatusBadRequest {
			t.Errorf("Handler returned a wrong status code: got %v but want %v", status, http.StatusBadRequest)
		}
	}
}

func TestCreateAccountWithInvalidRequest(t *testing.T) {

	test := []byte(`{"document_number": 81620264}`)

	req, _ := http.NewRequest("POST", "/account", bytes.NewBuffer(test))

	handler := http.HandlerFunc(accountController.CreateAccount)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code
	if status != http.StatusInternalServerError {
		t.Errorf("Handler returned a wrong status code: got %v but want %v", status, http.StatusInternalServerError)
	}
}

func TestCreateInvalidAccount(t *testing.T) {
	var testList = [][]byte{[]byte(`{"document_number": "706204590001739897"}`), []byte(`{"document_number": "81620264"}`),
		[]byte(`{"document_number": "flgne443"}`)}

	for _, test := range testList {
		req, _ := http.NewRequest("POST", "/account", bytes.NewBuffer(test))

		handler := http.HandlerFunc(accountController.CreateAccount)

		response := httptest.NewRecorder()

		handler.ServeHTTP(response, req)

		status := response.Code
		if status != http.StatusBadRequest {
			t.Errorf("Handler returned a wrong status code: got %v but want %v", status, http.StatusBadRequest)
		}
	}
}

func TestCreateAccount(t *testing.T) {
	var testList = [][]byte{[]byte(`{"document_number": "70620459000173"}`), []byte(`{"document_number": "81620264013"}`)}

	for _, test := range testList {
		req, _ := http.NewRequest("POST", "/account", bytes.NewBuffer(test))

		handler := http.HandlerFunc(accountController.CreateAccount)

		response := httptest.NewRecorder()

		handler.ServeHTTP(response, req)

		status := response.Code
		if status != http.StatusCreated {
			t.Errorf("Handler returned a wrong status code: got %v but want %v", status, http.StatusCreated)
		}

		var accountDTO domain.AccountDTO
		json.NewDecoder(io.Reader(response.Body)).Decode(&accountDTO)

		assert.NotNil(t, accountDTO.ID)
		validAccountID = accountDTO.ID
	}
}

func TestFindAccount(t *testing.T) {

	req, _ := http.NewRequest("GET", "/account", nil)

	handler := http.HandlerFunc(accountController.FindAccount)

	vars := map[string]string{
		"id": strconv.FormatInt(validAccountID, 10),
	}

	req = mux.SetURLVars(req, vars)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v but want %v", status, http.StatusOK)
	}
}

func TestCreatePositiveTransaction(t *testing.T) {
	var testList = []string{`{"account_id": ` + strconv.FormatInt(validAccountID, 10) + `,"operation_type_id": 4,"amount": 100.45}`,
		`{"account_id": ` + strconv.FormatInt(validAccountID, 10) + `,"operation_type_id": 4,"amount": 600.56}`}

	for _, test := range testList {
		var body = []byte(test)
		req, _ := http.NewRequest("POST", "/transaction", bytes.NewBuffer(body))

		handler := http.HandlerFunc(transactionController.CreateTransaction)

		response := httptest.NewRecorder()

		handler.ServeHTTP(response, req)

		status := response.Code
		if status != http.StatusCreated {
			t.Errorf("Handler returned a wrong status code: got %v but want %v", status, http.StatusCreated)
		}

		var transactionDTO domain.TransactionDTO
		json.NewDecoder(io.Reader(response.Body)).Decode(&transactionDTO)

		assert.NotNil(t, transactionDTO.ID)
		assert.Positive(t, transactionDTO.Amount)
	}
}

func TestCreateNegativeTransaction(t *testing.T) {
	var testList = []string{`{"account_id": ` + strconv.FormatInt(validAccountID, 10) + `,"operation_type_id": 1,"amount": 100.45}`,
		`{"account_id": ` + strconv.FormatInt(validAccountID, 10) + `,"operation_type_id": 2,"amount": 700.45}`,
		`{"account_id": ` + strconv.FormatInt(validAccountID, 10) + `,"operation_type_id": 3,"amount": 0.45}`}

	for _, test := range testList {
		var body = []byte(test)
		req, _ := http.NewRequest("POST", "/transaction", bytes.NewBuffer(body))

		handler := http.HandlerFunc(transactionController.CreateTransaction)

		response := httptest.NewRecorder()

		handler.ServeHTTP(response, req)

		status := response.Code
		if status != http.StatusCreated {
			t.Errorf("Handler returned a wrong status code: got %v but want %v", status, http.StatusCreated)
		}

		var transactionDTO domain.TransactionDTO
		json.NewDecoder(io.Reader(response.Body)).Decode(&transactionDTO)

		assert.NotNil(t, transactionDTO.ID)
		assert.Negative(t, transactionDTO.Amount)
	}
}

func TestCreateTransactionWithInvalidRequest(t *testing.T) {
	var testList = []string{`{"account_id": ` + strconv.FormatInt(validAccountID, 10) + `,"operation_type_id": 5,"amount": 100.45}`,
		`{"account_id": ` + strconv.FormatInt(validAccountID, 10) + `,"operation_type_id": 3,"amount": 0}`}

	for _, test := range testList {
		var body = []byte(test)
		req, _ := http.NewRequest("POST", "/transaction", bytes.NewBuffer(body))

		handler := http.HandlerFunc(transactionController.CreateTransaction)

		response := httptest.NewRecorder()

		handler.ServeHTTP(response, req)

		status := response.Code
		if status != http.StatusBadRequest {
			t.Errorf("Handler returned a wrong status code: got %v but want %v", status, http.StatusBadRequest)
		}
	}
}

func TestCreateInvalidTransaction(t *testing.T) {
	cleanDB()
	var testList = []string{`{"account_id": ` + strconv.FormatInt(validAccountID, 10) + `,"operation_type_id": 1,"amount": 100.45}`,
		`{"account_id": ` + strconv.FormatInt(validAccountID, 10) + `,"operation_type_id": "2","amount": 700.45}`}

	for _, test := range testList {
		var body = []byte(test)
		req, _ := http.NewRequest("POST", "/transaction", bytes.NewBuffer(body))

		handler := http.HandlerFunc(transactionController.CreateTransaction)

		response := httptest.NewRecorder()

		handler.ServeHTTP(response, req)

		status := response.Code
		if status != http.StatusInternalServerError {
			t.Errorf("Handler returned a wrong status code: got %v but want %v", status, http.StatusInternalServerError)
		}
	}
}

func cleanDB() {
	repo.ClearUp()
}
