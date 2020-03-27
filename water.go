package main

import (
	"fmt"
	"time"
)

type watering struct {
	ID                  int64 `gorm:"primary_key"`
	CreatedAt           time.Time
	Duration            time.Duration
	RawDataTicks        float64
	EventType           string //schedule, manual
	TotalVolumeInLiters float64
}

func (w *watering) Validate() error {
	return nil
}

// ** Ideal Scheduled Watering Flow **
// 1. Cron writes watering{EventType, Duration, StartTime} to db
// 2. Device gets ID, Duration, StartTime, EventType:
//   a. Battery Valve: On checkin
//	 b. Mains Valve: Device gets subscription to table on startup
// 3. Device waters
// 4. Device PUTS update for WaterVolume, Status of watering

// Add new water event to database
func (w *watering) SaveWatering() (*watering, error) {
	db.AutoMigrate(&watering{})
	newWatering := w
	// detect if manual or schedule, fill in.
	err := db.Create(&newWatering).Error
	if err != nil {
		return &watering{}, err
	}
	return newWatering, nil
}

// Update a watering event
func (w *watering) updateWatering() {
	// detect if manual or schedule
	db.AutoMigrate(&watering{})
	newWatering := w
	db.Create(&newWatering)
}

// Get water events from database
func getWaterFlowDatapoint() {
	newestWatering := watering{}
	db.Last(&newestWatering)
	// TODO: return something
	fmt.Printf("Newest water Flow: %v liters at %v for %vms\n", newestWatering.TotalVolumeInLiters, newestWatering.CreatedAt.String(), newestWatering.Duration)
}
