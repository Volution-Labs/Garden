package main

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func initDB() error {
	dataFilepath := "./data/sqlite3.db"
	if err := checkAndCreateFileIfNotExists(dataFilepath); err != nil {
		return err
	}

	var err error
	db, err = gorm.Open("sqlite3", dataFilepath)
	if err != nil {
		return err
	}

	db.LogMode(true)
	return nil
}

func checkAndCreateFileIfNotExists(filename string) error {
	if exists := doesFileExist(filename); !exists {
		if _, err := os.Create(filename); err != nil {
			return err
		}
	}
	return nil
}

// doesFileExists checks if a file exists and is not a
// directory before we try using it to prevent further errors.
func doesFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
