package dto

type CreateAccount struct {
	DocumentNumber string `json:"document_number"`
}

type Account struct {
	ID             int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type CreateTransaction struct {
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}
