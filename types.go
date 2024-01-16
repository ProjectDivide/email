package email

import "net/smtp"

type Email struct {
	FromAddress   string
	Body          string
	Authorization smtp.Auth
	SendAddress   string
}
