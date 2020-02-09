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
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Error loading .env file. Please see ./config/exampledotenv")
	}
	httpRouter := NewHttpRouter()
	coapRouter := NewCoapRouter()

	go func() {
		if err := http.ListenAndServe(os.Getenv("HTTP_SERVER_PORT"), httpRouter); err != nil {
			panic(err)
		}
	}()

	if err := coap.ListenAndServe(":5683", "udp", coapRouter); err != nil {
		panic(err)
	}
}
