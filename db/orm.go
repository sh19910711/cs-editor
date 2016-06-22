package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var ORM *gorm.DB

func Init() error {
	// setup database
	db, err := gorm.Open("sqlite3", ":memory:") // TODO: change driver
	if err != nil {
		return err
	}
	ORM = db
	return nil
}

func Close() {
	ORM.Close()
}
