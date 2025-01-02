// filepath: /c:/Go_Tutorial/main.go
package main

import (
	"log"
	
	initializer "webapp/initializer"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
    _ "webapp/docs"
	"webapp/controllers"
	
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

	

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// r.POST("/students", controllers.CreateStudent)
    r.GET("/students", controllers.GetStudents)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong1",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
