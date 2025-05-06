// filepath: /c:/Go_Tutorial/routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
	"net/http"
	"webapp/internal/api/controllers"
	"webapp/internal/api/middleware"
)

const (
	ADMIN_ROLE = "admin"
)

// RegisterRoutes sets up the API routes
func RegisterRoutes(r *gin.Engine,
	accountController controllers.AccountController,
	bookedSeatController controllers.BookedSeatController,
	bookingController controllers.BookingController,
	theaterController controllers.TheaterController,
	roomController controllers.RoomController,
	movieController controllers.MovieController,
	slotController controllers.SlotController,
	seatController controllers.SeatController,
	paymentController controllers.WebHookController,
	redisClient *redis.Client,
) {
	// Public Routes
	r.POST("/login", accountController.LoginHandler)
	r.POST("/register", accountController.RegisterHandler)

	accountGroup := r.Group("/v1/account")
	accountGroup.Use(middleware.JWTAuthMiddleware(redisClient), middleware.RoleMiddleware(ADMIN_ROLE))
	{
		accountGroup.GET("/list-all", accountController.ListAllAccountsHandler)
		accountGroup.GET("/random", accountController.GetRandomAccountHandler)
		accountGroup.GET("/:id", accountController.GetAccountByIDHandler)
		accountGroup.GET("/no-home", accountController.GetHomelessAccountsHandler)

	}

	theaterGroup := r.Group("/v1/theater")
	theaterGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	{
		theaterGroup.POST("/create", theaterController.CreateTheaterHandler)
		theaterGroup.GET("/list-all", theaterController.GetAllTheatersHandler)
	}

	roomGroup := r.Group("/v1/room")
	roomGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	{
		roomGroup.POST("/create", roomController.CreateRoomHandler)
	}

	slotGroup := r.Group("/v1/slot")
	slotGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	{
		slotGroup.POST("/create", slotController.CreateSlotHandler)
		slotGroup.GET("/list-all/:movieId", slotController.GetAllSlotsByMovieIdHandler)
	}

	movieGroup := r.Group("/v1/movie")
	movieGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	{
		movieGroup.GET("/:id", movieController.GetMovieByIDHandler)
		movieGroup.POST("/create", movieController.CreateMovieHandler)
		movieGroup.GET("/list-all", movieController.GetAllMoviesHandler)
	}

	bookingGroup := r.Group("/v1/booking")
	bookingGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	{
		bookingGroup.POST("/create", bookingController.CreateBookingHandler)
		bookingGroup.POST("/confirm/:bookingID", bookingController.ConfirmBookingHandler)
		bookingGroup.POST("/test", bookingController.TestingRabbitMq)
	}

	seatGroup := r.Group("/v1/seat")
	seatGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	{
		seatGroup.POST("/create", seatController.CreateSeatHandler)
	}

	paymentGroup := r.Group("/v1/payment")
	paymentGroup.POST("/webhook", paymentController.WebhookHandler)
	paymentGroup.POST("/create/:bookingId", paymentController.CreatePaymentLink)

	// Health Check
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
