package services

type MailServiceInterface interface {
	SendMail(to string, subject string, body string) error
	SendMailWithQueue() error
}
