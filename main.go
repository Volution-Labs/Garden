package main

import (
	"io"
	"log"
	"net/http"
	"os"

	coap "github.com/go-ocf/go-coap"
	"github.com/joho/godotenv"
)

func main() {
	// Logging setup
	f, err := os.OpenFile("output.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, f))

	// Load DB
	initDB()

	// Load .env
	err = godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Error loading .env file. Please see ./config/exampledotenv")
	}
	httpRouter := NewHttpRouter()
	coapRouter := NewCoapRouter()

	log.Println("Server started")
	go func() {
		if err := http.ListenAndServe(os.Getenv("HTTP_SERVER_PORT"), httpRouter); err != nil {
			panic(err)
		}
	}()
	if err := coap.ListenAndServe(":5683", "udp", coapRouter); err != nil {
		panic(err)
	}
}
