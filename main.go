// filepath: /c:/Go_Tutorial/main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	_ "webapp/docs"
	"webapp/initializer"
	"webapp/middleware"
	"webapp/routes"
	"webapp/services"
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
	config := &redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}
	var conn = *redis.NewClient(config)
	var addressService = services.NewAddressService(initializer.DB)

	var accountService = services.NewAccountService(initializer.DB, addressService, &conn)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(middleware.CORSMiddleware())
	routes.RegisterRoutes(r, accountService, &conn)

	err := r.Run()

	if err != nil {
		return
	}

}
