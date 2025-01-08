// filepath: /c:/Go_Tutorial/main.go
package main

import (
	"github.com/joho/godotenv"
	"log"
	_ "webapp/docs"
	"webapp/internal/infrastructure/cache"
	"webapp/internal/infrastructure/database"
)

func init() {
	err := godotenv.Load()
	database.ConnectDb()

	cache.ConnectRedis()

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

	//var addressService = services2.NewAddressService(database.DB)
	//var accountRepository = repositories.NewAccountRepository(database.DB)
	//var accountService = services2.NewAccountService(database.DB, addressService, cache.RedisClient, accountRepository)
	//
	//r := gin.Default()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	//
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//
	//r.Use(middleware.CORSMiddleware())
	//routes.RegisterRoutes(r, accountService, cache.RedisClient)
	//
	//err := r.Run()
	//
	//if err != nil {
	//	return
	//}

	r, err := InitializeApp()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	if err := r.Run(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
