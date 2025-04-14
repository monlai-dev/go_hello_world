// filepath: /c:/Go_Tutorial/main.go
package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"log"
	"os"
	"webapp/cmd/fx/accountfx"
	"webapp/cmd/fx/addressfx"
	"webapp/cmd/fx/bookedseatfx"
	"webapp/cmd/fx/bookingfx"
	"webapp/cmd/fx/cronjobfx"
	"webapp/cmd/fx/dbfx"
	"webapp/cmd/fx/mailfx"
	"webapp/cmd/fx/moviefx"
	"webapp/cmd/fx/paymentfx"
	"webapp/cmd/fx/rabbitmqfx"
	"webapp/cmd/fx/redisfx"
	"webapp/cmd/fx/roomfx"
	"webapp/cmd/fx/seatfx"
	"webapp/cmd/fx/slotfx"
	"webapp/cmd/fx/theaterfx"
	"webapp/cmd/fx/websocketfx"
	"webapp/internal/api/middleware"
	"webapp/internal/api/routes"
	"webapp/internal/infrastructure/cache"
	"webapp/internal/infrastructure/database"
	"webapp/internal/infrastructure/rabbitMq"
	"webapp/internal/services"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//database.ConnectDb()
	//cache.ConnectRedis()
	//database.DB.AutoMigrate(&models.Account{}, &models.Address{}, &models.Theater{}, &models.Movie{}, &models.Room{}, &models.Slot{}, &models.Seat{}, &models.BookedSeat{}, &models.Booking{})
	prometheus.MustRegister()
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {

	app := fx.New(
		// Register all modules here
		fx.Invoke(database.InitPostgres),
		fx.Invoke(cache.InitRedis),
		fx.Invoke(database.ConnectDb),
		fx.Invoke(cache.ConnectRedis),
		dbfx.Module,
		redisfx.Module,
		addressfx.Module,
		theaterfx.Module,
		roomfx.Module,
		accountfx.Module,
		moviefx.Module,
		slotfx.Module,
		seatfx.Module,
		bookedseatfx.Module,
		cronjobfx.Module,
		rabbitmqfx.Module,
		mailfx.Module,
		bookingfx.Module,
		paymentfx.Module,
		websocketfx.Module,

		// Register your router
		fx.Provide(ProvideRouter),

		// Start the HTTP server
		fx.Invoke(StartServer),
		fx.Invoke(StartCronJob),
		fx.Invoke(ConsumeMail),
	)

	app.Run()

}

func StartServer(lc fx.Lifecycle, engine *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Starting HTTP server at ${PORT}")
				if err := engine.Run(":" + os.Getenv("PORT")); err != nil {
					log.Fatalf("Failed to start server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping HTTP server")
			return nil
		},
	})
}

func StartCronJob(lc fx.Lifecycle, cronJobService *services.CronJobService, bookingServiceInterface services.BookingServiceInterface) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Starting cron job")

			_, err := cronJobService.AddFunc("@every 1m", func() {
				log.Printf("Running scheduler")
				err := bookingServiceInterface.Scheduler()
				if err != nil {
					log.Printf("Error while running scheduler: %v", err)
				}
			})

			if err != nil {
				log.Printf("Error while adding cron job: %v", err)
			}
			cronJobService.StartCronJob()
			// Add cron job
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping cron job")

			return nil
		},
	})
}

func ConsumeMail(lc fx.Lifecycle, mailService services.MailServiceInterface) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Starting mail consumer")
			if consumeErr := mailService.SendMailWithQueue(); consumeErr != nil {
				log.Fatalf("Error consuming message: %v", consumeErr)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping mail consumer")
			return nil
		},
	})
}

func ProvideRouter(
	accountService services.AccountServiceInterface,
	redisClient *redis.Client,
	roomServiceInterface services.RoomServiceInterface,
	theaterServiceInterface services.TheaterServiceInterface,
	slotServiceInterface services.SlotServiceInterface,
	movieService services.MovieServiceInterface,
	bookingServiceInterface services.BookingServiceInterface,
	seatServiceInterface services.SeatServiceInterface,
	paymentService services.PaymentServiceInterface,
	socketService *services.WebsocketService,
	rabbitClient *rabbitMq.RabbitMq,
) *gin.Engine {
	log.Println("ProvideRouter called, initializing gin.Engine")
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RateLimitMiddleware())
	socketService.AttachToRouter(r)

	routes.RegisterRoutes(r,
		accountService,
		redisClient,
		roomServiceInterface,
		theaterServiceInterface,
		slotServiceInterface,
		movieService,
		bookingServiceInterface,
		seatServiceInterface,
		paymentService,
		socketService,
		rabbitClient)
	return r
}
