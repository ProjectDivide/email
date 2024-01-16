package email

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func NewEmail() Email {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	gmailPassword := os.Getenv("GMAIL_PASSWORD")
	smtpEndpoint := os.Getenv("SMTP_ENDPOINT")
	port := os.Getenv("SMTP_PORT")
	FromAddress := os.Getenv("FROM_ADDRESS")

	auth := smtp.PlainAuth("", FromAddress, gmailPassword, smtpEndpoint)

	return Email{
		FromAddress:   FromAddress,
		Authorization: auth,
		Body:          "",
		SendAddress:   smtpEndpoint + ":" + port,
	}
}

func (e Email) Send(toaddress string, body []byte) bool {

	toList := []string{toaddress}

	Must(
		smtp.SendMail(
			e.SendAddress,
			e.Authorization,
			e.FromAddress,
			toList,
			body,
		),
		"Error sending email",
	)
	return true

}
