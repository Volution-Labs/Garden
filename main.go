package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Load DB
	initDB()

	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
