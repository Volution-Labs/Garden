package main

import (
	"log"
	"os"
	"time"

	coap "github.com/go-ocf/go-coap"
	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func handleA(w coap.ResponseWriter, req *coap.Request) {
	log.Printf("Got message in handleA: path=%q: %#v from %v", req.Msg.Path(), req.Msg, req.Client.RemoteAddr())
	w.SetContentFormat(coap.TextPlain)
	log.Printf("Transmitting from A")
	if _, err := w.Write([]byte("hello world")); err != nil {
		log.Printf("Cannot send response: %v", err)
	}
}

func handleB(w coap.ResponseWriter, req *coap.Request) {
	log.Printf("Got message in handleB: path=%q: %#v from %v", req.Msg.Path(), req.Msg, req.Client.RemoteAddr())
	resp := w.NewResponse(coap.Content)
	resp.SetOption(coap.ContentFormat, coap.TextPlain)
	resp.SetPayload([]byte("good bye!"))
	log.Printf("Transmitting from B %#v", resp)
	if err := w.WriteMsg(resp); err != nil {
		log.Printf("Cannot send response: %v", err)
	}
}

type Temp struct {
	TimeStamp *time.Time
	Value     int64
	Location  string
}

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Use .env
	appEnv := os.Getenv("APP_ENV")
	log.Printf(appEnv)

	// Use db
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("Error opening db")
	}
	defer db.Close()
	db.AutoMigrate(&Temp{})
	temp := Temp{Value: 18, Location: "lettuce"}
	db.Create(&temp)
	var read Temp
	db.Last(&read)
	log.Printf(read.Location)

	// Coap
	mux := coap.NewServeMux()
	mux.Handle("/a", coap.HandlerFunc(handleA))
	mux.Handle("/b", coap.HandlerFunc(handleB))

	coap.ListenAndServe(":5683", "udp", mux)
}
