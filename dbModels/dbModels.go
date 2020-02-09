package dbModels

import (
	"time"
)

// this needs workshopped...create foreign-key restraint instead??
type relation struct {
	Name     string
	Type     string
	Location string
}

type SoilTemp struct {
	ID            int64 `gorm:"primary_key"`
	CreatedAt     time.Time
	TempInCelsius float64
	Relationship  relation
}

type Moisture struct {
	ID                   int64 `gorm:"primary_key"`
	CreatedAt            time.Time
	PercentOfRelMoisture float64
	Relationship         relation
}

type Light struct {
	ID           int64 `gorm:"primary_key"`
	CreatedAt    time.Time
	Lux          int
	Relationship relation
}

type WaterFlow struct {
	ID                  int64 `gorm:"primary_key"`
	CreatedAt           time.Time
	Duration            time.Time
	TotalVolumeInLiters float64
	Relationship        relation
}
