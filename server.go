package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()

	const port string = ":8000"

	router.HandleFunc("/", func(resp http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(resp, "Server up and running...")
	})
	router.HandleFunc("/account", getAccount).Methods("GET")
	router.HandleFunc("/account", createAccount).Methods("POST")
	router.HandleFunc("/transaction", createTransaction).Methods("POST")

	log.Println("server listining on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
