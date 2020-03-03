package notification

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Sends a push notification via Pushover service
func SendPushover(m Message) {
	message := map[string]interface{}{
		"token":    os.Getenv("PUSHOVER_TOKEN"),
		"user":     os.Getenv("PUSHOVER_USER"),
		"message":  m.Message,
		"priority": m.Priority,
	}

	jsonPayload, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://api.pushover.net/1/messages.json", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])
}
