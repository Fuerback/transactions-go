package main

import (
	"fmt"
	"net/http"

	"github.com/Fuerback/transactions-go/controller"
	"github.com/Fuerback/transactions-go/router"
)

var (
	httpRouter            router.Router                    = router.NewMuxRouter()
	transactionController controller.TransactionController = controller.NewTransactionController()
	accountController     controller.AccountController     = controller.NewAccountController()
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
