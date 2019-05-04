package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("Error opening db")
	}
	db.LogMode(true)
}
