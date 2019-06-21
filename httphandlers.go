package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
)

// API: Get and return specified amount of data points
func getChartData(w http.ResponseWriter, r *http.Request) {
	listOfTemps := SoilTemp{}
	db.Find(&listOfTemps)
	fmt.Printf("Newest Temperature: %v\u2103C at %v\n", listOfTemps.Temperature, listOfTemps.CreatedAt.String())
}

// API: Set to water on next update for lenght of time or turn off.
func ManualWater(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	fmt.Fprintln(w, "Cool! Will work on that")
}

// API:
func sendTemp(w http.ResponseWriter, r *http.Request) {
	//getTemps()
	fmt.Fprintln(w, "Here ya go:")
}
