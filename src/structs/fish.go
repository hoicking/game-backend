package structs


type Fish struct{
	ID int32 `gorm:"type:int;PRIMARY KEY;AUTO INCREMENT;NOT NULL"`
	Name string `gorm:"type:varchar(40);NOT NULL"`
	Alias string `gorm:"type:varchar(40);DEFAULT NULL"` 
	Max int32 `gorm:"type:int;DEFAULT NULL"`
	Min int32 `gorm:"type:int;DEFAULT NULL"`
	Unit int32 `gorm:"type:int;DEFAULT NULL"`
	Style int32 `gorm:"type:int;DEFAULT NULL"`
	Probability int32 `gorm:"type:int;DEFAULT NULL"`
	Position string `gorm:"type:varchar(20);DEFAULT NULL"` 
	Description string `gorm:"type:varchar(255);DEFAULT NULL"` 
}