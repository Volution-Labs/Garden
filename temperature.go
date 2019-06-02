package main

import (
	"log"
	"time"
)

// add new temp to db
func newTemp(temp float64, measurementType string, location string) {
	db.AutoMigrate(&soilTemp{})
	newTemp := soilTemp{TimeStamp: time.Now(), Value: temp, Location: location}
	db.Create(&newTemp)
}

// get
func getTemps(dateTimeStart time.Time, dataTimeEnd time.Time) {
	lastTemp := soilTemp{}
	db.Last(&lastTemp)
	log.Printf(lastTemp.Location)
}
