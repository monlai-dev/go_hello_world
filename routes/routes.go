// filepath: /c:/Go_Tutorial/routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"webapp/middleware"
	"webapp/models/request_models"
	"webapp/models/response_models"
	"webapp/services"
)

func RegisterRoutes(r *gin.Engine, accountService services.AccountServiceInterface) {
	// Public routes
	// @Summary Login
	// @Description Login
	// @Tags account
	// @Accept json
	// @Produce json
	// @Param email body string true "Email"
	// @Param password body string true "Password"
	r.POST("/login", func(c *gin.Context) {
		var loginRequest request_models.LoginRequest
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := accountService.Login(loginRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response_models.Response{
				ResponseCode: http.StatusForbidden,
				Message:      "Invalid email or password",
			})
			return
		}

		c.JSON(http.StatusOK, response_models.Response{
			ResponseCode: http.StatusOK,
			Message:      "Login successful",
			Data:         []interface{}{token},
		})
	})

	r.POST("/register", func(c *gin.Context) {
		var account request_models.RegisterRequest

		if err := c.ShouldBindJSON(&account); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		createdAccount, err := accountService.CreateAccount(account)

		if err != nil {
			c.JSON(http.StatusInternalServerError, response_models.Response{
				ResponseCode: http.StatusInternalServerError,
				Message:      "Error creating account",
			})
			return
		}

		c.JSON(http.StatusOK, response_models.Response{
			ResponseCode: http.StatusOK,
			Message:      "Account created successfully",
			Data:         []interface{}{createdAccount},
		})
	})

	// Protected routes
	accountGroup := r.Group("/account")
	accountGroup.Use(middleware.JWTAuthMiddleware())
	{
		// @Summary Get all accounts
		// @Description Get all accounts
		// @Tags account
		// @Accept json
		// @Produce json
		// @Success 200 {object} []models.Account
		// @Router /account/list-all [get]
		accountGroup.GET("/list-all", func(c *gin.Context) {
			accounts, err := accountService.GetAllAccounts()
			if err != nil {
				c.JSON(http.StatusOK, response_models.Response{
					ResponseCode: http.StatusBadRequest,
					Message:      "Error getting accounts",
				})
				return
			}

			c.JSON(http.StatusOK, response_models.Response{
				ResponseCode: http.StatusOK,
				Message:      "Accounts retrieved successfully",
				Data:         []interface{}{accounts},
			})
		})

		accountGroup.GET("/random", func(c *gin.Context) {
			account, err := accountService.GetRandomAccount()
			if err != nil {
				c.JSON(http.StatusOK, response_models.Response{
					ResponseCode: http.StatusBadRequest,
					Message:      "Error getting account",
				})
				return
			}
			c.JSON(http.StatusOK, response_models.Response{
				ResponseCode: http.StatusOK,
				Message:      "Account retrieved successfully",
				Data:         []interface{}{account},
			})
		})

		accountGroup.GET("/:id", func(c *gin.Context) {
			id, _ := strconv.Atoi(c.Param("id"))
			account, err := accountService.GetAccountById(id)

			if err != nil {
				c.JSON(http.StatusOK, response_models.Response{
					ResponseCode: http.StatusBadRequest,
					Message:      "Error getting account",
				})
				return
			}

			c.JSON(http.StatusOK, response_models.Response{
				ResponseCode: http.StatusOK,
				Message:      "Account retrieved successfully",
				Data:         []interface{}{account},
			})
		})

		accountGroup.GET("/no-home", func(c *gin.Context) {
			account, err := accountService.GetAllHomelessAccounts()

			if err != nil {
				c.JSON(http.StatusOK, response_models.Response{
					ResponseCode: http.StatusBadRequest,
					Message:      "Error",
				})
				return
			}

			c.JSON(http.StatusOK, response_models.Response{
				ResponseCode: http.StatusOK,
				Message:      "Account retrieved successfully",
				Data:         []interface{}{account},
			})

		})

		accountGroup.PUT("/:email", func(c *gin.Context) {
			email := c.Param("email")
			var address request_models.AddressRequest
			if err := c.ShouldBindJSON(&address); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			err := accountService.UpdateAddress(email, address)

			if err != nil {
				c.JSON(http.StatusInternalServerError, response_models.Response{
					ResponseCode: http.StatusInternalServerError,
					Message:      "Error updating address",
				})
				return
			}

			c.JSON(http.StatusOK, response_models.Response{
				ResponseCode: http.StatusOK,
				Message:      "Address updated successfully",
			})
		})
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
