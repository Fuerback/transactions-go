package errors

import "errors"

var ErrInvalidOperation = errors.New("Invalid operation")

type ServiceError struct {
	Message string `json:"message"`
}
