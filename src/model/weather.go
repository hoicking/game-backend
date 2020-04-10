package model

import (
	"game-backend/src/structs"
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"time"
)

type WeatherModel struct {
	DB *gorm.DB
}


func (model *WeatherModel) SaveWeather(data string) {

	fmtDatda := structs.Weather{
		ID: ksuid.New().String(),
		Weather: data,
		Date: time.Now(),
	}
	model.DB.Table("weather").Create(&fmtDatda)
}