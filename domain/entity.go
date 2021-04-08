package domain

import (
	"github.com/Fuerback/transactions-go/errors"
)

type Account struct {
	ID             int64
	DocumentNumber string
}

type Transaction struct {
	ID            int
	AccountID     int
	OperationType int
	Amount        float64
	EventDate     string
}

const (
	UNKNOWN = iota
	VISTA
	PARCELADA
	SAQUE
	PAGAMENTO
)

func IsNegative(i int) (bool, error) {
	switch i {
	case VISTA:
		return true, nil
	case PARCELADA:
		return true, nil
	case SAQUE:
		return true, nil
	case PAGAMENTO:
		return false, nil
	default:
		return false, errors.ErrInvalidOperation
	}
}
