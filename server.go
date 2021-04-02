package main

import (
	"fmt"
	"net/http"

	"github.com/Fuerback/transactions-go/controller"
	"github.com/Fuerback/transactions-go/router"
	"github.com/Fuerback/transactions-go/service"
)

var (
	// criar repository aqui e passar pros serviços
	httpRouter            router.Router                    = router.NewMuxRouter()
	transactionService    service.TransactionService       = service.NewTransactionService()
	accountService        service.AccountService           = service.NewAccountService()
	transactionController controller.TransactionController = controller.NewTransactionController(transactionService)
	accountController     controller.AccountController     = controller.NewAccountController(accountService)
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(resp, "Server up and running...")
	})
	httpRouter.GET("/account", accountController.FindAccount)
	httpRouter.POST("/account", accountController.CreateAccount)
	httpRouter.POST("/transaction", transactionController.CreateTransaction)

	httpRouter.SERVE(port)
}
