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
func RegisterRoutes(r *gin.Engine, accountService services.AccountServiceInterface, redisClient *redis.Client) {
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

	// Health Check
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
