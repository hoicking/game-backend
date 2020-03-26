package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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
func regist(r *mux.Router) {
	r.HandleFunc("/info/weather", getWeather).Methods("get")
	r.HandleFunc("/info/data", getData).Methods("get")

}
