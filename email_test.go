package email

import "testing"

func TestSend(t *testing.T) {
	email := NewEmail()
	if email.Send("gulegulzaradnan@gmail.com", []byte("Email form go test")) == false {
		t.Error("Cannot send test email")
	}
}
