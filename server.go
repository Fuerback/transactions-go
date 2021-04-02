package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Fuerback/transactions-go/controller"
	"github.com/Fuerback/transactions-go/repository"
	"github.com/Fuerback/transactions-go/router"
	"github.com/Fuerback/transactions-go/service"
	_ "github.com/mattn/go-sqlite3"
)

const port string = ":8000"

func main() {
	db, err := sql.Open("sqlite3", "db/transaction.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqliteRepo := repository.NewSqlite(db)

	httpRouter := router.NewMuxRouter()

	transactionService := service.NewTransactionService(sqliteRepo)
	accountService := service.NewAccountService(sqliteRepo)

	transactionController := controller.NewTransactionController(transactionService)
	accountController := controller.NewAccountController(accountService)

	httpRouter.GET("/", func(resp http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(resp, "Server up and running...")
	})
	httpRouter.GET("/account/{id}", accountController.FindAccount)
	httpRouter.POST("/account", accountController.CreateAccount)
	httpRouter.POST("/transaction", transactionController.CreateTransaction)

	httpRouter.SERVE(port)
}
