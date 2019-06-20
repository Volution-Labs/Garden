package main

import (
	"fmt"
	"time"
)

type SoilTemp struct {
	ID          int64 `gorm:"primary_key"`
	CreatedAt   time.Time
	Temperature float64
}

// Add new temperature to database
func newSoilTempDatapoint(temperature float64) {
	if temperature > -90 && temperature < 150 {
		db.AutoMigrate(&SoilTemp{})
		newTempReading := SoilTemp{CreatedAt: time.Now(), Temperature: temperature}
		db.Create(&newTempReading)
	}
}

// Get temperature(s) from database
func getTempDatapoints(numberOfPoints int) {
	newestTemp := SoilTemp{}
	db.Last(&newestTemp)
	// Return something but print for now.
	fmt.Printf("Newest Temperature: %v\u2103C at %v\n", newestTemp.Temperature, newestTemp.CreatedAt.String())
}

// Get
func getTempDatapoint(time time.Time) {
	newestTemp := SoilTemp{}
	db.Last(&newestTemp)
	// Return something but print for now.
	fmt.Printf("Newest Temperature: %v\u2103C at %v\n", newestTemp.Temperature, newestTemp.CreatedAt.String())
}
