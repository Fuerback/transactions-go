package entity

import (
	"github.com/Fuerback/transactions-go/errors"
)

type Account struct {
	ID             int
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

func GetOperation(i int) (string, error) {
	switch i {
	case VISTA:
		return "VISTA", nil
	case PARCELADA:
		return "PARCELADA", nil
	case SAQUE:
		return "SAQUE", nil
	case PAGAMENTO:
		return "PAGAMENTO", nil
	default:
		return "", errors.ErrInvalidOperation
	}
}
