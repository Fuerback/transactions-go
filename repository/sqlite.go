package repository

import (
	"database/sql"
	"time"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/entity"
)

type sqlite struct {
	DB *sql.DB
}

func (s *sqlite) CreateAccount(account *dto.CreateAccount) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into account(document_number) values (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(account.DocumentNumber)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *sqlite) FindAccount(ID int64) (entity.Account, error) {
	u := entity.Account{}

	stmt, err := s.DB.Prepare("select * from account where id = ?")
	if err != nil {
		return entity.Account{}, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&u.ID, &u.DocumentNumber)
	if err != nil {
		return entity.Account{}, err
	}

	return u, nil
}

func (s *sqlite) CreateTransaction(transaction *dto.CreateTransaction) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into transaction(account_id, amount, event_date, operation_type) values (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(transaction.AccountID, transaction.Amount, time.Now(), transaction.OperationTypeID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
