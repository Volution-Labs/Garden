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
	if err := initDB(); err != nil {
		log.Fatal(err)
	}
	log.Println("✨  DB started successfully  ✨")

	// Load .env
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Fatal("Error loading .env file: %s", err)
	}
	log.Println("✨  config loaded successfully  ✨")

	// Create routers
	httpRouter := NewHttpRouter()
	coapRouter := NewCoapRouter()
	log.Println("✨  routers created  ✨")

	// Start coap
	go func() {
		if err := coap.ListenAndServe(os.Getenv("COAP_SERVER_PORT"), "udp", coapRouter); err != nil {
			log.Fatal(err)
		}
	}()
	log.Printf("\n\ncoap started successfully, listening at: %s", os.Getenv("COAP_SERVER_PORT"))

	// Start app
	log.Printf("\n\napp started successfully, listening at: %s", os.Getenv("HTTP_SERVER_PORT"))
	if err := http.ListenAndServe(os.Getenv("HTTP_SERVER_PORT"), httpRouter); err != nil {
		log.Fatal(err)
	}
}
