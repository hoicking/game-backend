package controller

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type InfoCtl struct {
	Router *mux.Router
}

func getWeather(res http.ResponseWriter, request *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	data := map[string]string{"aa": "bb"}

	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	res.WriteHeader(http.StatusOK)
	res.Write(json)
}

func getData(res http.ResponseWriter, request *http.Request) {

}

// InitRouter sfdsf
func (ctl InfoCtl) Regist() {
	ctl.Router.HandleFunc("/info/weather", getWeather).Methods("get")
	ctl.Router.HandleFunc("/info/data", getData).Methods("get")
}
