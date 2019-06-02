package main

import (
	"time"
)

type baseModel struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
}

// this needs workshopped...create foreign-key restraint instead??
type relation struct {
	Name     string
	Type     string
	Location string
}

type soilTemp struct {
	baseModel
	TempInCelsius float64
	relation
}

type moisture struct {
	baseModel
	PercentOfRelMoisture float64
	relation
}

type light struct {
	baseModel
	Lux int
	relation
}

type waterFlow struct {
	baseModel
	Duration            time.Time
	TotalVolumeInLiters float64
	relation
}
