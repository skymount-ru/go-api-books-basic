package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var (
	DBConn *gorm.DB
)

type Book struct {
	gorm.Model
	Name string
	Country string
	Released string
}

func OpenDB() *gorm.DB {
	var err error
	var DSN = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_BASE"),
	)
	DBConn, err = gorm.Open("mysql", DSN)
	if err != nil {
		panic("There was an error opening the database! " + err.Error())
	}
	log.Println("Connected to database successfully!")
	DBConn.AutoMigrate(&Book{})
	return DBConn
}
