package model

import (
	"game-backend/src/structs"
	"github.com/segmentio/ksuid"
	"time"
)

func (model *BaseModel) SaveWeather(data string) {

	fmtDatda := structs.Weather{
		ID: ksuid.New().String(),
		Weather: data,
		Date: time.Now(),
	}
	model.db.Table("weather").Create(&fmtDatda)
}