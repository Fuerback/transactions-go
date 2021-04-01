package dto

type CreateAccount struct {
	DocumentNumber string `json:"document_number"`
}

type GetAccount struct {
	ID             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type CreateTransaction struct {
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}
