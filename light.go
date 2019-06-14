package main

import (
	"fmt"
	"time"
)

type light struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	Lux       float64
}

// Add new light to database
func newLightDatapoint(lightReading float64) {
	if lightReading >= 0 && lightReading < 33000 {
		db.AutoMigrate(&light{})
		newLightReading := light{CreatedAt: time.Now(), Lux: lightReading}
		db.Create(&newLightReading)
	}
}

// Get light(s) from database
func getLightDatapoint() {
	newestLight := light{}
	db.Last(&newestLight)
	// Return something but print for now.
	fmt.Printf("Newest light reading: %v lux at %v\n", newestLight.Lux, newestLight.CreatedAt.String())
}
