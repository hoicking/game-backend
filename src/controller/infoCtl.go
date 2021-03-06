package controller

import (
	"encoding/json"
	"net/http"
	"game-backend/src/model"

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

func saveWeather(res http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var data map[string]string
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	weatherModel := model.NewBaseModel()

	weatherModel.SaveWeather(data["weather"])
	res.Write([]byte("success"))

}

// InitRouter
func (ctl InfoCtl) Regist() {
	ctl.Router.HandleFunc("/info/weather", getWeather).Methods("get")
	ctl.Router.HandleFunc("/info/weather", saveWeather).Methods("put")
}
