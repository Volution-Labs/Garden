package main

import (
	coap "github.com/go-ocf/go-coap"
)

func NewCoapRouter() *coap.ServeMux {
	router := coap.NewServeMux()
	router.Handle("/sensors", coap.HandlerFunc(handleSensorData))
	//router.Handle("/messages", coap.HandlerFunc(handlemessages))
	return router
}
