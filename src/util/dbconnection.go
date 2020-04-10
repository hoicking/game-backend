package util

import (
	// "fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() (err error) {

	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MSYQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connection := user + ":" + pass + "@(" + host + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"

	DB, err = gorm.Open("mysql", connection)

	if err != nil {
		return nil
	}

	err = DB.DB().Ping()


	if err != nil {
		return err
	}

	DB.LogMode(true)
	return nil
}
