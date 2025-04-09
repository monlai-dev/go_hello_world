package mailfx

import (
	"go.uber.org/fx"
	gomail "gopkg.in/mail.v2"
	"os"
	"webapp/internal/infrastructure/rabbitMq"
	"webapp/internal/services"
)

var Module = fx.Provide(provideMailService)

func provideMailService(rabbitmqClient *rabbitMq.RabbitMq) services.MailServiceInterface {
	return services.NewMailService(gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD")), rabbitmqClient)
}
