// filepath: /c:/Go_Tutorial/controllers/accountControllers.go
package controllers

import (
	"log"
	"net/http"
	"webapp/initializer"
	models "webapp/models/db_models"
	"github.com/gin-gonic/gin"
)

	


func CreateStudent(c *gin.Context) {

	account := c.ShouldBind(&models.Account{})

	result := initializer.DB.Create(&account)

	if condition := result.Error; condition != nil {
		log.Fatal(condition)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when creating student"})
	}

	c.JSON(http.StatusOK, account)

}


func GetStudents(c *gin.Context) {

	var accounts []models.Account
	initializer.DB.Find(&accounts)
	c.JSON(http.StatusOK, accounts)
}


func DeleteStudent(c *gin.Context) {

	result := initializer.DB.Where("id = ?", c.Param("id")).Delete(&models.Account{})
	
	if condition := result.Error; condition != nil {
		log.Fatal(condition)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when deleting student"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}
