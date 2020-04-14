package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"game-backend/src/model"
	"encoding/json"

)

type FishCtl struct {
	Router *mux.Router
}

func getFishes(res http.ResponseWriter, request *http.Request) {
	fishModel := model.NewBaseModel()

	fishes := fishModel.GetFishes()

	res.Header().Set("Content-Type", "application/json")

	json, err := json.Marshal(fishes)
	if err != nil {
		panic(err)
	}
	res.WriteHeader(http.StatusOK)
	res.Write(json)

}

func (ctl FishCtl) Regist() {
	ctl.Router.HandleFunc("/fish/fishes", getFishes).Methods("get")
}
