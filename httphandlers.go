package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	coap "github.com/go-ocf/go-coap"
	"github.com/gorilla/mux"
)

// API: Get and return specified amount of data points
func getChartData(w http.ResponseWriter, r *http.Request) {
	var temps []SoilTemp
	db.Order("id desc").Limit(20).Find(&temps)
	json.NewEncoder(w).Encode(&temps)
	//var temps []SoilTemp
	db.Order("id desc").Limit(20).Find(&temps)
	json.NewEncoder(w).Encode(&temps)
}

// API: Set to water on next update for lenght of time or turn off.
func ManualWater(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	state := vars["key"]
	co, err := coap.Dial("udp", "garden.local:5683")
	if err != nil {
		fmt.Printf("Error dialing: %v", err)
	}
	resp, err := co.Post("/water", coap.MediaType(50), strings.NewReader(state))
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
	}
	if string(resp.Payload()) == "ok" {
		fmt.Printf("Valve set to %v\n", state)
	} else {
		fmt.Printf("Error while setting %v | Response: %v\n", state, string(resp.Payload()))
	}
}

// API:
func sendTemp(w http.ResponseWriter, r *http.Request) {
	//getTemps()
	fmt.Fprintln(w, "Here ya go:")
}
