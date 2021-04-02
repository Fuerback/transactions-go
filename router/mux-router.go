package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var muxDispatcher = mux.NewRouter()

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (r *muxRouter) GET(uri string, f func(resp http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (r *muxRouter) POST(uri string, f func(resp http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (r *muxRouter) SERVE(port string) {
	fmt.Printf("Mux running on port %v\n", port)
	http.ListenAndServe(port, muxDispatcher)
}
