package main

import (
	"fmt"
	"time"
)

type waterFlow struct {
	ID                  int64 `gorm:"primary_key"`
	CreatedAt           time.Time
	Duration            time.Time
	TotalVolumeInLiters float64
}

// Add new flow to database
func newWaterFlowDatapoint(flowReading float64, duration time.Time) {
	if flowReading > 0 && flowReading < 900 {
		db.AutoMigrate(&waterFlow{})
		newFlowReading := waterFlow{CreatedAt: time.Now(), TotalVolumeInLiters: flowReading, Duration: duration}
		db.Create(&newFlowReading)
	}
}

// Get flow(s) from database
func getWaterFlowDatapoint() {
	newestWater := waterFlow{}
	db.Last(&newestWater)
	// Return something but print for now.
	fmt.Printf("Newest water Flow: %v liters at %v for %vms\n", newestWater.TotalVolumeInLiters, newestWater.CreatedAt.String(), newestWater.Duration)
}
