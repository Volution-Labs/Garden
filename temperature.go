package main

import (
	"fmt"
	"time"

	"github.com/Volution-Labs/garden-server/dbModels"
)

// type SoilTemp struct {
// 	ID          int64 `gorm:"primary_key"`
// 	CreatedAt   time.Time
// 	Temperature float64
// }

// Add new temperature to database
func newSoilTempDatapoint(temp float64) {
	if temp > -90 && temp < 150 {
		db.AutoMigrate(&dbModels.SoilTemp{})

		newTempReading := dbModels.SoilTemp{
			CreatedAt:     time.Now(),
			TempInCelsius: temp,
		}

		db.Create(&newTempReading)
	}
}

// Get temperature(s) from database
func getTempDatapoints(numberOfPoints int) {
	newestTemp := dbModels.SoilTemp{}
	db.Last(&newestTemp)

	// Return something but print for now.
	fmt.Printf("Newest Temperature: %v\u2103C at %v\n", newestTemp.TempInCelsius, newestTemp.CreatedAt.String())
}

// Get
func getTempDatapoint(time time.Time) {
	newestTemp := dbModels.SoilTemp{}
	db.Last(&newestTemp)

	// Return something but print for now.
	fmt.Printf("Newest Temperature: %v\u2103C at %v\n", newestTemp.TempInCelsius, newestTemp.CreatedAt.String())
}
