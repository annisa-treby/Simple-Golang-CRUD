package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

func Config() *gorm.DB {
	conn := os.Getenv("URL")
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}