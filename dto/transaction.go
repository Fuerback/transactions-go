package dto

type CreateAccount struct {
	DocumentNumber string `json:"document_number" validate:"required"`
}

type Account struct {
	ID             int64  `json:"account_id" validate:"required"`
	DocumentNumber string `json:"document_number" validate:"required"`
}

type CreateTransaction struct {
	AccountID       int64   `json:"account_id" validate:"required"`
	OperationTypeID int     `json:"operation_type_id" validate:"required"`
	Amount          float64 `json:"amount" validate:"required"`
}

type Transaction struct {
	ID              int64
	AccountID       int64
	OperationTypeID int
	Amount          float64
}
