package model

import (
	"game-backend/src/structs"
	"github.com/jinzhu/gorm"

)

type WeatherModel struct {
	db *gorm.DB
}

func (model *WeatherModel) SaveWeather(data string) {

	fmtDatda := structs.Weather{
		Weather: data,
	}
	model.db.Table("weather").Create(&fmtDatda)
}