package main

import (
	"fmt"
	"time"
)

type moisture struct {
	ID                   int64 `gorm:"primary_key"`
	CreatedAt            time.Time
	PercentOfRelMoisture float64
}

// Add new moisture to database
func newMoistureDatapoint(moistureReading float64) {
	if moistureReading >= 0 && moistureReading <= 100 {
		db.AutoMigrate(&moisture{})
		newMoistureReading := moisture{CreatedAt: time.Now(), PercentOfRelMoisture: moistureReading}
		db.Create(&newMoistureReading)
	}
}

// Get moisture(s) from database
func getMoistureDatapoint() {
	newestMoisture := moisture{}
	db.Last(&newestMoisture)
	// Return something but print for now.
	fmt.Printf("Newest Moisture: %v%% at %v\n", newestMoisture.PercentOfRelMoisture, newestMoisture.CreatedAt.String())
}
