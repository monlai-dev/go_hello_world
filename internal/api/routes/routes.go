// filepath: /c:/Go_Tutorial/routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"webapp/internal/api/controllers"
)

const (
	ADMIN_ROLE = "admin"
)

// RegisterRoutes sets up the API routes
func RegisterRoutes(r *gin.Engine,
	accountController controllers.AccountController) {
	// Public Routes
	r.POST("/login", accountController.LoginHandler)
	r.POST("/register", accountController.RegisterHandler)

	// Protected Routes
	//accountGroup := r.Group("/v1/account")
	//accountGroup.Use(middleware.JWTAuthMiddleware(redisClient), middleware.RoleMiddleware(ADMIN_ROLE))
	//{
	//	accountGroup.GET("/list-all", controllers.ListAllAccountsHandler(accountService))
	//	accountGroup.GET("/random", controllers.GetRandomAccountHandler(accountService))
	//	accountGroup.GET("/:id", controllers.GetAccountByIDHandler(accountService))
	//	accountGroup.GET("/no-home", controllers.GetHomelessAccountsHandler(accountService))
	//	accountGroup.PUT("/update-address", controllers.UpdateAddressHandler(accountService))
	//	accountGroup.POST("/logout", controllers.LogoutHandler(accountService))
	//}
	//
	//theaterGroup := r.Group("/v1/theater")
	//theaterGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	//{
	//	theaterGroup.POST("/create", controllers.CreateTheaterHandler(theaterService))
	//	theaterGroup.GET("/list-all", controllers.GetAllTheatersHandler(theaterService))
	//}
	//
	//roomGroup := r.Group("/v1/room")
	//roomGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	//{
	//	roomGroup.POST("/create", controllers.CreateRoomHandler(roomService))
	//}
	//
	//slotGroup := r.Group("/v1/slot")
	//slotGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	//{
	//	slotGroup.POST("/create", controllers.CreateSlotHandler(slotService))
	//	slotGroup.GET("/list-all/:movieId", controllers.GetAllSlotsByMovieIdHandler(slotService))
	//}
	//
	//movieGroup := r.Group("/v1/movie")
	//movieGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	//{
	//	movieGroup.GET("/:id", controllers.GetMovieByIDHandler(movieService))
	//	movieGroup.POST("/create", controllers.CreateMovieHandler(movieService))
	//	movieGroup.GET("/list-all", controllers.GetAllMoviesHandler(movieService))
	//}
	//
	//bookingGroup := r.Group("/v1/booking")
	//bookingGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	//{
	//	bookingGroup.POST("/create", controllers.CreateBookingHandler(bookingService, rabbitClient))
	//	bookingGroup.POST("/confirm/:bookingID", controllers.ConfirmBookingHandler(bookingService))
	//	bookingGroup.POST("/test", controllers.TestingRabbitMq(bookingService))
	//}
	//
	//seatGroup := r.Group("/v1/seat")
	//seatGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	//{
	//	seatGroup.POST("/create", controllers.CreateSeatHandler(seatService))
	//}
	//
	//paymentGroup := r.Group("/v1/payment")
	//paymentGroup.POST("/webhook", controllers.WebhookHandler(bookingService))
	//paymentGroup.POST("/create/:bookingId", controllers.CreatePaymentLink(paymentService))

	// Health Check
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
