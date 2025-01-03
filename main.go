// filepath: /c:/Go_Tutorial/main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	_ "webapp/docs"
	initializer "webapp/initializer"
	"webapp/routes"
	"webapp/services"
	"webapp/middleware"
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

	r.Use(middleware.CORSMiddleware())
	routes.RegisterRoutes(r, accountService)
	
	r.Run() // listen and serve on 0.0.0.0:8080

}
