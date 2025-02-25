package services

import (
	"context"
	"errors"
	"github.com/goccy/go-json"
	amqp "github.com/rabbitmq/amqp091-go"
	gomail "gopkg.in/mail.v2"
	"log"
	"os"
	"strconv"
	"sync"
	"webapp/internal/infrastructure/rabbitMq"
)

type EmailTask struct {
	EmailRequest EmailRequest
	Delivery     amqp.Delivery
}

type MailService struct {
	dialer      *gomail.Dialer
	rabbit      *rabbitMq.RabbitMq
	taskChannel chan EmailTask
	bufferPool  int
	wg          *sync.WaitGroup
	ctx         context.Context // Added for lifecycle management
	cancel      context.CancelFunc
}

func NewMailService(dialer *gomail.Dialer, rabbitClient *rabbitMq.RabbitMq) MailServiceInterface {
	workers, _ := strconv.Atoi(os.Getenv("WORKERS_POOL"))
	ctx, cancel := context.WithCancel(context.Background())

	ms := &MailService{
		taskChannel: make(chan EmailTask, workers),
		dialer:      dialer,
		rabbit:      rabbitClient,
		bufferPool:  workers,
		wg:          &sync.WaitGroup{},
		ctx:         ctx,
		cancel:      cancel,
	}
	go func() {
		ms.startWorkers()
	}()

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
	// only 5 workers are allowed to process the email tasks concurrently
	for i := 0; i < 5; i++ {
		m.wg.Add(1)
		go m.worker(i)
	}
	m.wg.Wait()
	log.Printf("All workers have stopped")
}

func (m MailService) worker(workerId int) {

	defer m.wg.Done()

	for {
		select {
		case <-m.ctx.Done():
			log.Printf("Worker %d stopped due to context cancellation", workerId)
			return
		case task, ok := <-m.taskChannel:
			if !ok {
				log.Printf("Worker %d stopped: task channel closed", workerId)
				return
			}
			// Process the email task
			err := m.SendMail(
				task.EmailRequest.Email,
				task.EmailRequest.Subject,
				task.EmailRequest.Body,
			)
			if err != nil {
				log.Printf("Worker %d failed to send email to %s: %v", workerId, task.EmailRequest.Email, err)
				_ = task.Delivery.Nack(false, true)
			} else {
				log.Printf("Worker %d sent email to %s", workerId, task.EmailRequest.Email)
				_ = task.Delivery.Ack(false)
			}
		}
	}
}
