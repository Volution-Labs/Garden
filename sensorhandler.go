package main

import (
	"encoding/json"
	"fmt"
	"log"

	coap "github.com/go-ocf/go-coap"
)

// Handle incoming sensor data.
// st - Soil Temp, l - Light, f - Flow (duration), m - moisture, at - Air Temp, r - Rain
func handleSensorData(w coap.ResponseWriter, req *coap.Request) {
	p := req.Msg.Payload()
	log.Printf("| New sensor data: %v\n", string(p))
	var dat map[string]interface{}
	if err := json.Unmarshal(p, &dat); err != nil {
		log.Println("handleSensorData: Could not convert unmarshal payload")
	}
	for k, v := range dat {
		switch k {
		case "st":
			fmt.Printf("Soil Temp: %.2f\n", v)
		case "l":
			fmt.Printf("Light: %.0f\n", v)
		case "f":
			fmt.Printf("Flow: %.0f\n", v)
		case "m":
			fmt.Printf("Moisture: %.0f\n", v)
		default:
			log.Printf("Err: Unknown datapoint in sensor payload: {\"%v\": %v}\n", k, v)
		}
	}
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Printf("Cannot send response: %v", err)
	}
}
