package main

import (
	"fmt"
	"game-backend/src/controller"
	"game-backend/src/util"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	err := util.InitDB()

	if err != nil {
		log.Fatal("init db fail: ", err)
	}

	port := os.Getenv("PORT")

	fmt.Println("start...",port)
	router := mux.NewRouter().StrictSlash(true)
	controller.InfoCtl{Router: router}.Regist()

	fmtPort := ":" + port
	fmt.Println("ListenAndServe: ",fmtPort)


	if err := http.ListenAndServe(fmtPort, router); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
