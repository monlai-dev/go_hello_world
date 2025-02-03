package services

import (
	"errors"
	gomail "gopkg.in/mail.v2"
	"log"
)

type MailService struct {
	dialer *gomail.Dialer
}

func NewMailService(dialer *gomail.Dialer) MailServiceInterface {
	return &MailService{
		dialer: dialer,
	}
}

func (m MailService) SendMail(to string, subject string, body string) error {

	mail := gomail.NewMessage()
	mail.SetHeader("From", "")
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	err := m.dialer.DialAndSend(mail)
	if err != nil {
		log.Printf("Error sending mail: %v", err)
		return errors.New("error sending mail")
	}

	return nil
}
