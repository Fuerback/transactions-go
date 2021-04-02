package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/errors"
	"github.com/Fuerback/transactions-go/service"
)

var accountService service.AccountService = service.NewAccountService()

type AccountController interface {
	CreateAccount(resp http.ResponseWriter, r *http.Request)
	FindAccount(resp http.ResponseWriter, r *http.Request)
}

type accountController struct{}

func NewAccountController() AccountController {
	return &accountController{}
}

func (a *accountController) CreateAccount(resp http.ResponseWriter, r *http.Request) {
	accountDTO := new(dto.CreateAccount)
	err := json.NewDecoder(r.Body).Decode(accountDTO)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error unmarshalling the request"}`))
		return
	}
	err = accountService.Create(accountDTO)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	resp.WriteHeader(http.StatusCreated)
}

func (a *accountController) FindAccount(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	accountDTO := new(dto.FindAccount)
	err := json.NewDecoder(r.Body).Decode(accountDTO)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error unmarshalling the request"}`))
		return
	}
	accounts, err := accountService.Find(accountDTO)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	result, err := json.Marshal(accounts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error unmarshalling the accounts"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}
