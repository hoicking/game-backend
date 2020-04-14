package model

import (
	"game-backend/src/structs"
	"fmt"
)

func (model *BaseModel) GetFishes() []structs.Fish {
	var fishes []structs.Fish
	model.db.Table("fish").Select("*").Find(&fishes)

	fmt.Println("---------",fishes)
	return fishes
}