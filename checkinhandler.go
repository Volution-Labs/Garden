package main

import (
	"fmt"
	"log"

	coap "github.com/go-ocf/go-coap"
)

// Handle checkins from battery powered/normaly sleeping devices
func handleCheckin(w coap.ResponseWriter, req *coap.Request) {
	message := req.Msg.Payload()
	fmt.Printf("New device checkin: %v\n", string(message))
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Printf("Cannot send response: %v", err)
	}
}
