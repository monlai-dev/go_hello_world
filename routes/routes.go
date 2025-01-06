// filepath: /c:/Go_Tutorial/routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
	"webapp/middleware"
	"webapp/models/request_models"
	"webapp/models/response_models"
	"webapp/services"
)

// RegisterRoutes sets up the API routes
func RegisterRoutes(r *gin.Engine, accountService services.AccountServiceInterface, redisClient *redis.Client) {
	// Public Routes
	r.POST("/login", loginHandler(accountService))
	r.POST("/register", registerHandler(accountService))

	// Protected Routes
	accountGroup := r.Group("/account")
	accountGroup.Use(middleware.JWTAuthMiddleware(redisClient))
	{
		accountGroup.GET("/list-all", listAllAccountsHandler(accountService))
		accountGroup.GET("/random", getRandomAccountHandler(accountService))
		accountGroup.GET("/:id", getAccountByIDHandler(accountService))
		accountGroup.GET("/no-home", getHomelessAccountsHandler(accountService))
		accountGroup.PUT("/update-address", updateAddressHandler(accountService))
		accountGroup.POST("/logout", logoutHandler(accountService))
	}

	// Health Check
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}

// Public Handlers
func loginHandler(accountService services.AccountServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request_models.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := accountService.Login(req)
		if err != nil {
			c.JSON(http.StatusForbidden, responseError("Invalid email or password"))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Login successful", []interface{}{gin.H{"token": token}}))
	}
}

func registerHandler(accountService services.AccountServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request_models.RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		createdAccount, err := accountService.CreateAccount(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, responseSuccess("Account created successfully", []interface{}{createdAccount}))
	}
}

// Protected Handlers
func listAllAccountsHandler(accountService services.AccountServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {

		page, _ := strconv.Atoi(c.Query("page"))
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))

		accounts, err := accountService.GetAllAccounts(page, pageSize)

		if err != nil {
			c.JSON(http.StatusBadRequest, responseError("Error getting accounts"))
			return
		}
		c.JSON(http.StatusOK, responseSuccess("Accounts retrieved successfully", []interface{}{accounts}))
	}
}

func getRandomAccountHandler(accountService services.AccountServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		account, err := accountService.GetRandomAccount()
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError("Error getting account"))
			return
		}
		c.JSON(http.StatusOK, responseSuccess("Account retrieved successfully", []interface{}{account}))
	}
}

func getAccountByIDHandler(accountService services.AccountServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		account, err := accountService.GetAccountById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError("Error getting account"))
			return
		}
		c.JSON(http.StatusOK, responseSuccess("Account retrieved successfully", []interface{}{account}))
	}
}

func getHomelessAccountsHandler(accountService services.AccountServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		accounts, err := accountService.GetAllHomelessAccounts()
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError("Error retrieving accounts"))
			return
		}
		c.JSON(http.StatusOK, responseSuccess("Accounts retrieved successfully", []interface{}{accounts}))
	}
}

func updateAddressHandler(accountService services.AccountServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		email, _ := c.Get("email")
		var address request_models.AddressRequest
		if err := c.ShouldBindJSON(&address); err != nil {
			c.JSON(http.StatusBadRequest, responseError("Invalid address"))
			return
		}

		err := accountService.UpdateAddress(email.(string), address)
		if err != nil {
			c.JSON(http.StatusBadRequest, responseError("Error updating address"))
			return
		}
		c.JSON(http.StatusOK, responseSuccess("Address updated successfully", nil))
	}
}

func logoutHandler(accountService services.AccountServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token := authHeader[7:] // Extract token from "Bearer <token>"

		if err := accountService.Logout(token); err != nil {
			c.JSON(http.StatusBadRequest, responseError("Error logging out"))
			return
		}
		c.JSON(http.StatusOK, responseSuccess("Logged out successfully", nil))
	}
}

// Response Helpers
func responseError(message string) response_models.Response {
	return response_models.Response{
		ResponseCode: http.StatusBadRequest,
		Message:      message,
	}
}

func responseSuccess(message string, data []interface{}) response_models.Response {
	return response_models.Response{
		ResponseCode: http.StatusOK,
		Message:      message,
		Data:         data,
	}
}
