package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DBCon *gorm.DB
)

func InitDB(filepath string) {
	var err error
	DBCon, err = gorm.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if DBCon == nil {
		panic("Failed to connect to database.")
	}
}

