package model

import (

)

func (model * BaseModel) GetFises() {

	model.db.Table("fish").Select()

}