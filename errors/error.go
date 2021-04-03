package errors

import "errors"

var ErrInvalidOperation = errors.New("Invalid operation")
var ErrInvalidDocumentNumber = errors.New("Invalid document number")

type ServiceError struct {
	Message string `json:"message"`
}
