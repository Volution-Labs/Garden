package main

import (
	"encoding/json"
	"log"
	"time"

	coap "github.com/go-ocf/go-coap"
)

func handleWateringComplete(w coap.ResponseWriter, req *coap.Request) {
	// Payload: f = flow in gallons, fraw = raw flow ticks, d = water duration, id = schedule id
	p := req.Msg.Payload()
	var jsonPayload map[string]interface{}
	if err := json.Unmarshal(p, &jsonPayload); err != nil {
		log.Println("handleWaterFinished: Could not unmarshal payload")
		w.Write([]byte("ERR"))
		return
	}
	log.Printf("Watering Completed: Duration = %v, Flow = %v", jsonPayload["d"], jsonPayload["fraw"])

	dataset := watering{}
	for k, v := range jsonPayload {
		switch k {
		case "f":
			dataset.TotalVolumeInLiters = jsonPayload["f"].(float64)
		case "d":
			dataset.Duration = time.Duration(jsonPayload["d"].(float64)) * time.Millisecond
		case "fraw":
			dataset.RawDataTicks = jsonPayload["fraw"].(float64)
		case "id":
			dataset.ID = jsonPayload["ID"].(int64)
		default:
			log.Printf("Err: Unknown datapoint in water-complete payload: {\"%v\": %v}\n", k, v)
		}
	}
	dataset.SaveWatering()

	if _, err := w.Write([]byte("OK")); err != nil {
		log.Printf("Cannot send response: %v", err)
	}
}
