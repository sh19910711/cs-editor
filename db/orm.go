package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var ORM *gorm.DB

func Init() error {
	dburl := os.Getenv("DATABASE_URL")
	if dburl == "" {
		dburl = ":memory:"
	}
	db, err := gorm.Open("sqlite3", dburl) // TODO: change driver
	if err != nil {
		return err
	}
	ORM = db
	return nil
}

func Close() {
	ORM.Close()
}
