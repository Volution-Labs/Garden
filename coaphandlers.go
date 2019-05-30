package main

import (
	"fmt"
	"log"

	coap "github.com/go-ocf/go-coap"
)

func handleSensorData(w coap.ResponseWriter, req *coap.Request) {
	log.Printf("Got message in handleA: path=%q: %#v from %v", req.Msg.Path(), req.Msg, req.Client.RemoteAddr())
	message := req.Msg.Payload()
	fmt.Printf(string(message))
	// tf, err := strconv.ParseFloat(ts, 32)
	// if err != nil {
	// 	panic("Error opening db")
	// }
	// newTemp(tf, "soil", loc)
	// fmt.Fprintln(w, "Added to database:", tf)
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Printf("Cannot send response: %v", err)
	}
}
