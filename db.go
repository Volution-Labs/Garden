package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open("sqlite3", "./data/sqlite3.db")
	if err != nil {
		panic("Error opening db")
	}
	db.LogMode(true)
}
