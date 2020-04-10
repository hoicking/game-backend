package model
import (
	"github.com/jinzhu/gorm"
	"game-backend/src/util"
)

type BaseModel struct{
	db *gorm.DB
}

func NewBaseModel() BaseModel{
	model := BaseModel{
		db: util.DB,
	}

	return model
}