package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Println(e.Error())
	}

	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DBNAME")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbDriver := "mysql"

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)

	conn, err := gorm.Open(dbDriver, dbURI)

	if err != nil {
		fmt.Println(err.Error())
	}

	db = conn
	db.Debug().AutoMigrate(&Activity{}, &Todo{})
}

func GetDB() *gorm.DB {
	return db
}
