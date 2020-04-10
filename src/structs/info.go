package structs

import (
	"time"
)

//Weather for table weather
type Weather struct{
	ID string `gorm:"type:varchar(50)"`
	Date  time.Time `gorm:"type:date;NOT NULL"`
	Weather string 	`gorm:"type:varchar(20);NOT NULL"`
}