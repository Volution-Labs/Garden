package notification

import (
	"log"
	"net/smtp"
	"os"
)

// Sends a notification via Email
func SendMail(m *Message) {
	from := os.Getenv("EMAIL_USER")
	pass := os.Getenv("EMAIL_PASSWORD")
	to := os.Getenv("EMAIL_TO")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + m.Subject + "\n\n" +
		m.Message

	err := smtp.SendMail(os.Getenv("EMAIL_SMTP")+os.Getenv("EMAIL_PORT"),
		smtp.PlainAuth("", from, pass, os.Getenv("EMAIL_SMTP")),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("Email sent: " + m.Message)
}
