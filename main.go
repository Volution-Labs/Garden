package main

import (
	"log"
	"net/http"
	"os"

	coap "github.com/go-ocf/go-coap"
	"github.com/joho/godotenv"
)

func main() {
	// Load DB
	initDB()

	// Load .env
	err := godotenv.Load("../config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	httpRouter := NewHttpRouter()
	coapRouter := NewCoapRouter()

	go func() {
		http.ListenAndServe(os.Getenv("HTTP_SERVER_PORT"), httpRouter)
	}()
	coap.ListenAndServe(":5683", "udp", coapRouter)
}
