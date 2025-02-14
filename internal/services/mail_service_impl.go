package services

import (
	"errors"
	"github.com/goccy/go-json"
	gomail "gopkg.in/mail.v2"
	"log"
	"webapp/internal/infrastructure/rabbitMq"
)

type MailService struct {
	dialer *gomail.Dialer
	rabbit *rabbitMq.RabbitMq
}

func NewMailService(dialer *gomail.Dialer, rabbitClient *rabbitMq.RabbitMq) MailServiceInterface {
	return &MailService{
		dialer: dialer,
		rabbit: rabbitClient,
	}
}

func (m MailService) SendMail(to string, subject string, body string) error {

	mail := gomail.NewMessage()
	mail.SetHeader("From", "maihailongviet@gmail.com")
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

func (m MailService) SendMailWithQueue() error {

	message, err := m.rabbit.Consume("email_queue", "email_service")

	if err != nil {
		log.Printf("Error consuming message: %v", err)
		return errors.New("error consuming message")
	}

	go func() {
		for d := range message {
			var emailMessage EmailRequest
			if err := json.Unmarshal(d.Body, &emailMessage); err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}

			if err := m.SendMail(emailMessage.Email, emailMessage.Subject, emailMessage.Body); err != nil {
				log.Println("Error sending email: ", err)
			}

			err := d.Ack(true)
			if err != nil {
				return
			}
		}
	}()
	return nil
}
