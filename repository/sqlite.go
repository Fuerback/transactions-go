package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/Fuerback/transactions-go/dto"
	"github.com/Fuerback/transactions-go/entity"
)

type sqlite struct {
	DBFilePath string
}

func (s *sqlite) CreateAccount(account *dto.CreateAccount) (int64, error) {
	db, err := sql.Open("sqlite3", s.DBFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	stmt, err := tx.Prepare("insert into account(document_number) values (?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(account.DocumentNumber)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	ID, _ := result.LastInsertId()
	tx.Commit()
	return ID, nil
}

func (s *sqlite) FindAccount(ID int64) (entity.Account, error) {
	db, err := sql.Open("sqlite3", s.DBFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	u := entity.Account{}

	stmt, err := db.Prepare("select * from account where id = ?")
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

func (s *sqlite) CreateTransaction(transaction *dto.CreateTransaction) (int64, error) {
	db, err := sql.Open("sqlite3", s.DBFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	stmt, err := tx.Prepare("insert into [transaction](account_id, amount, event_date, operation_type) values (?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(transaction.AccountID, transaction.Amount, time.Now().Format(time.RFC3339), transaction.OperationTypeID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	ID, _ := result.LastInsertId()
	tx.Commit()
	return ID, nil
}

func (s *sqlite) ClearUp() error {
	db, err := sql.Open("sqlite3", s.DBFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from [transaction]")
	_, err = tx.Exec("delete from account")
	tx.Commit()
	return err
}
