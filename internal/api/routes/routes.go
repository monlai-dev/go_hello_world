// filepath: /c:/Go_Tutorial/routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"webapp/internal/api/controllers"
	"webapp/internal/api/middleware"
	"webapp/internal/services"
)

const (
	ADMIN_ROLE = "admin"
)

// RegisterRoutes sets up the API routes
func RegisterRoutes(r *gin.Engine, accountService services.AccountServiceInterface, redisClient *redis.Client, roomService services.RoomServiceInterface, theaterService services.TheaterServiceInterface, slotService services.SlotServiceInterface, movieService services.MovieServiceInterface) {
	// Public Routes
	r.POST("/login", controllers.LoginHandler(accountService))
	r.POST("/register", controllers.RegisterHandler(accountService))

	// Protected Routes
	accountGroup := r.Group("/v1/account")
	accountGroup.Use(middleware.JWTAuthMiddleware(redisClient), middleware.RoleMiddleware(ADMIN_ROLE))
	{
		accountGroup.GET("/list-all", controllers.ListAllAccountsHandler(accountService))
		accountGroup.GET("/random", controllers.GetRandomAccountHandler(accountService))
		accountGroup.GET("/:id", controllers.GetAccountByIDHandler(accountService))
		accountGroup.GET("/no-home", controllers.GetHomelessAccountsHandler(accountService))
		accountGroup.PUT("/update-address", controllers.UpdateAddressHandler(accountService))
		accountGroup.POST("/logout", controllers.LogoutHandler(accountService))
	}

	theaterGroup := r.Group("/v1/theater")
	theaterGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	{
		theaterGroup.POST("/create", controllers.CreateTheaterHandler(theaterService))
		theaterGroup.GET("/list-all", controllers.GetAllTheatersHandler(theaterService))
	}

	roomGroup := r.Group("/v1/room")
	roomGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	{
		roomGroup.POST("/create", controllers.CreateRoomHandler(roomService))
	}

	slotGroup := r.Group("/v1/slot")
	slotGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	{
		slotGroup.POST("/create", controllers.CreateSlotHandler(slotService))
		slotGroup.GET("/list-all/:movieId", controllers.GetAllSlotsByMovieIdHandler(slotService))
	}

	movieGroup := r.Group("/v1/movie")
	movieGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	{
		movieGroup.POST("/create", controllers.CreateMovieHandler(movieService))
	}

	// Health Check
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
