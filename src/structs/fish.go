package structs


type Fish struct{
	ID int32 `gorm:"type:int;PRIMARY KEY;AUTO INCREMENT;NOT NULL" json:"id"`
	Name string `gorm:"type:varchar(40);NOT NULL" json:"name"`
	Alias string `gorm:"type:varchar(40);DEFAULT NULL" json:"alias"` 
	Max int32 `gorm:"type:int;DEFAULT NULL" json:"max"`
	Min int32 `gorm:"type:int;DEFAULT NULL" json:"min"`
	Unit int32 `gorm:"type:int;DEFAULT NULL" json:"unit"`
	Style int32 `gorm:"type:int;DEFAULT NULL" json:"style"`
	Probability int32 `gorm:"type:int;DEFAULT NULL" json:"probability"`
	Position string `gorm:"type:varchar(20);DEFAULT NULL" json:"position"` 
	Description string `gorm:"type:varchar(255);DEFAULT NULL" json:"description"` 
}