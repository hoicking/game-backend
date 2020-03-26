package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc('/',)

	http.ListenAndServe(":8080", myRouter)
}
