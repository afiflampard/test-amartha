package helper

import (
	"log"
	"os"

	"gopkg.in/mail.v2"
)

func SendEmail(toEmail string, subject string, body string) error {
	// Set up email client
	m := mail.NewMessage()

	// Set from, to, subject, and body email
	m.SetHeader("From", os.Getenv("USER_EMAIL"))
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Setup SMTP client
	log.Println(os.Getenv("USER_EMAIL"))
	log.Println(os.Getenv("USER_PASSWORD"))
	log.Println(toEmail)
	d := mail.NewDialer("smtp.gmail.com", 465, os.Getenv("USER_EMAIL"), os.Getenv("USER_PASSWORD"))

	// Kirim email
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}
