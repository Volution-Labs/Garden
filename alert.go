package main

import "time"

type alert struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	Method    []string
	Subject   string
	Message   string
	Priority  int
}

func (a *alert) SaveAlert() (*alert, error) {
	db.AutoMigrate(&alert{})
	newAlert := a
	err := db.Create(&newAlert).Error
	if err != nil {
		return &alert{}, err
	}
	return newAlert, nil
}
