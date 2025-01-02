// filepath: /c:/Go_Tutorial/main.go
package main

import (
	"log"

	initializer "webapp/initializer"
	"webapp/middleware"
	"webapp/models/request_models"

	_ "webapp/docs"

	"webapp/services"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func init() {
	err := godotenv.Load()

	initializer.ConnectDb()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
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

	var accountService services.AccountService = services.NewAccountService(initializer.DB)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// r.POST("/students", controllers.CreateStudent)
	Group := r.Group("/account")
	Group.Use(middleware.JWTAuthMiddleware())

	Group.GET("/list-all", func(c *gin.Context) {

		accounts, err := accountService.GetAllAccounts()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, accounts)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong1",
		})
	})

	r.POST("/login", func(c *gin.Context) {

		var loginRequest request_models.LoginRequest

		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := accountService.Login(loginRequest)

		if err != nil {
			c.JSON(500, gin.H{"error": "Invalid email or password"})
			return
		}

		c.JSON(200, token)
	})

	r.POST("/register", func(c *gin.Context) {

		var account request_models.RegisterRequest
		
        if err := c.ShouldBindJSON(&account); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

		createdAccount, err := accountService.CreateAccount(account)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			log.Fatal(err)
			return
		}

		c.JSON(200, gin.H{"account": createdAccount})
	})

	r.Run() // listen and serve on 0.0.0.0:8080

}
