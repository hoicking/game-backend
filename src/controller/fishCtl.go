package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

type FishCtl struct {
	Router *mux.Router
}

func getFishes(res http.ResponseWriter, request *http.Request) {

}

func (ctl FishCtl) Regist() {
	ctl.Router.HandleFunc("/fish/fishes", getFishes).Methods("get")
}
