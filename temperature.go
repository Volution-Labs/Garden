package main

import (
	"log"
	"time"
)

type Temp struct {
	ID        int64 `gorm:"primary_key"`
	TimeStamp time.Time
	Value     float64
	Type      string
	Location  string
}

// add new temp to db
func newTemp(temp float64, measurementType string, location string) {
	db.AutoMigrate(&Temp{})
	newTemp := Temp{TimeStamp: time.Now(), Value: temp, Location: location}
	db.Create(&newTemp)
}

// get
func getTemps(dateTimeStart time.Time, dataTimeEnd time.Time) {
	lastTemp := Temp{}
	db.Last(&lastTemp)
	log.Printf(lastTemp.Location)
}
