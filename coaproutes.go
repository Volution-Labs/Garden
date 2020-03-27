package main

import (
	coap "github.com/go-ocf/go-coap"
)

func NewCoapRouter() *coap.ServeMux {
	router := coap.NewServeMux()
	router.Handle("/sensors", coap.HandlerFunc(handleSensorData))
	router.Handle("/checkin", coap.HandlerFunc(handleCheckin))
	router.Handle("/alert", coap.HandlerFunc(handleAlert))
	router.Handle("/watering-complete", coap.HandlerFunc(handleWateringComplete))
	return router
}
