package domain

type CreateAccountDTO struct {
	DocumentNumber string `json:"document_number" validate:"required,min=11,max=14"`
}

type AccountDTO struct {
	ID             int64  `json:"account_id" validate:"required"`
	DocumentNumber string `json:"document_number" validate:"required"`
}

type CreateTransactionDTO struct {
	AccountID       int64   `json:"account_id" validate:"required"`
	OperationTypeID int     `json:"operation_type_id" validate:"required"`
	Amount          float64 `json:"amount" validate:"required"`
}

type TransactionDTO struct {
	ID              int64
	AccountID       int64
	OperationTypeID int
	Amount          float64
}
