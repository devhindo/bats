package main

import (
	"log"
	"os"

	"github.com/resend/resend-go/v2"

)
func sendMailResend(to string, subject string, body string) error {
	apiKey := os.Getenv("RESEND")
	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    os.Getenv("RESEND_MAIL"),
		To: 	[]string{to},
		Subject: subject,
		Html:   body,
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		log.Println("Error sending mail: ", err)
		return err
	}

	log.Println("Mail sent successfully", sent)

	return nil
}
