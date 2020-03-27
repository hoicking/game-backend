package main

import (
	"game-backend/src/controller"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	controller.InfoCtl{Router: router}.Regist()
	http.ListenAndServe(":8080", router)
}
