package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/errors"
	"github.com/Fuerback/transactions-go/service"
	"github.com/gorilla/mux"
)

var accountService service.AccountService

type AccountController interface {
	CreateAccount(resp http.ResponseWriter, r *http.Request)
	FindAccount(resp http.ResponseWriter, r *http.Request)
}

type accountController struct{}

func NewAccountController(service service.AccountService) AccountController {
	accountService = service
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
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	resp.Header().Set("Content-type", "application/json")
	accounts, err := accountService.Find(ID)
	if err != nil {
		if err == sql.ErrNoRows {
			resp.WriteHeader(http.StatusBadRequest)
		} else {
			resp.WriteHeader(http.StatusInternalServerError)
		}
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
