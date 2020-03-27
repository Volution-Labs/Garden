package main

import (
	"encoding/json"
	"log"

	coap "github.com/go-ocf/go-coap"
)

// Handle/split incoming sensor data.
// st - Soil Temp, l - Light, f - Flow (duration), m - moisture, at - Air Temp, r - Rain
func handleSensorData(w coap.ResponseWriter, req *coap.Request) {
	p := req.Msg.Payload()
	log.Printf("New sensor data: %v\n", string(p))
	var jsonPayload map[string]interface{}
	if err := json.Unmarshal(p, &jsonPayload); err != nil {
		log.Println("handleSensorData: Could not convert unmarshal payload")
		w.Write([]byte("ERR"))
		return
	}
	for k, v := range jsonPayload {
		switch k {
		case "st":
			newSoilTempDatapoint(v.(float64))
		case "l":
			newLightDatapoint(v.(float64))
		case "f":
			// TODO: log extra flow somewhere
		case "m":
			newMoistureDatapoint(v.(float64))
		default:
			log.Printf("Err: Unknown datapoint in sensor payload: {\"%v\": %v}\n", k, v)
		}
	}
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Printf("Cannot send response: %v", err)
	}
}

// Handle emergancy alert from device
// json format: {"event", "message"}, Message types: extra water, system crash, low batteries,
func handleAlert(w coap.ResponseWriter, req *coap.Request) {
	p := req.Msg.Payload()
	log.Printf("| New sensor data: %v\n", string(p))
	var jsonPayload map[string]interface{}
	if err := json.Unmarshal(p, &jsonPayload); err != nil {
		log.Println("handleSensorData: Could not convert unmarshal payload")
		w.Write([]byte("ERR"))
		return
	}
	for k, v := range jsonPayload {
		switch k {
		case "st":
			newSoilTempDatapoint(v.(float64))
		case "l":
			newLightDatapoint(v.(float64))
		case "m":
			newMoistureDatapoint(v.(float64))
		default:
			log.Printf("Err: Unknown datapoint in sensor payload: {\"%v\": %v}\n", k, v)
		}
	}
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Printf("Cannot send response: %v", err)
	}
}
