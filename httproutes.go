package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewHttpRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	// Handle Statics
	router.
		PathPrefix("/").
		Handler(http.FileServer(http.Dir("./web/")))

	return router
}

var routes = Routes{
	Route{
		"SensorDataList",
		"GET",
		"/api/sensors",
		getChartData,
	},
	Route{
		"ManualWater",
		"POST",
		"/api/water/{key}",
		ManualWater,
	},
	Route{
		"sendTemp",
		"GET",
		"/api/temp/",
		sendTemp,
	},
}
