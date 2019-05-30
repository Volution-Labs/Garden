package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// API: Get and return sensor data
func ListData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The soil is on fire!!!!")
}

// API: Set to water on next update for lenght of time
func ManualWater(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	fmt.Fprintln(w, "Cool! Will work on that")
}

// API: Add temp
func addTemp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loc := vars["loc"]
	ts := vars["id"]
	tf, err := strconv.ParseFloat(ts, 32)
	if err != nil {
		panic("Error opening db")
	}
	newTemp(tf, "soil", loc)
	fmt.Fprintln(w, "Added to database:", tf)
}

// API:
func sendTemp(w http.ResponseWriter, r *http.Request) {
	//getTemps()
	fmt.Fprintln(w, "Here ya go:")
}
