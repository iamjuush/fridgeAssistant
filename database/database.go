package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DBCon *gorm.DB
)

func InitDB() {
	var err error
	DBCon, err = gorm.Open("mysql", "root:4Admin2Use@(localhost)/dbo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	if DBCon == nil {
		panic("Failed to connect to database.")
	}
}

