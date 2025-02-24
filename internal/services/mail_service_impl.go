package services

import (
	"errors"
	"github.com/goccy/go-json"
	amqp "github.com/rabbitmq/amqp091-go"
	gomail "gopkg.in/mail.v2"
	"log"
	"webapp/internal/infrastructure/rabbitMq"
)

const (
	WORKERS = 5
)

type EmailTask struct {
	EmailRequest EmailRequest
	Delivery     amqp.Delivery
}

type MailService struct {
	dialer      *gomail.Dialer
	rabbit      *rabbitMq.RabbitMq
	taskChannel chan EmailTask
}

func NewMailService(dialer *gomail.Dialer, rabbitClient *rabbitMq.RabbitMq) MailServiceInterface {
	ms := &MailService{
		taskChannel: make(chan EmailTask, WORKERS),
		dialer:      dialer,
		rabbit:      rabbitClient,
	}
	ms.startWorkers()
	return ms
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

			m.taskChannel <- EmailTask{
				EmailRequest: emailMessage,
				Delivery:     d,
			}
		}
	}()
	return nil
}

func (m MailService) startWorkers() {
	for i := 0; i < WORKERS; i++ {
		go m.worker(i)
	}
}

func (m MailService) worker(workerId int) {
	for task := range m.taskChannel {
		// Process the email task
		err := m.SendMail(
			task.EmailRequest.Email,
			task.EmailRequest.Subject,
			task.EmailRequest.Body,
		)
		if err != nil {
			log.Printf("Failed to send email to %s: %v", task.EmailRequest.Email, err)
			// Requeue the task on failure (optional)
			_ = task.Delivery.Nack(false, true)
		} else {
			// Acknowledge the message on success
			log.Printf("Worker %d sent email to %s", workerId, task.EmailRequest.Email)
			_ = task.Delivery.Ack(false)
		}
	}
}
