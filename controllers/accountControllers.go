// filepath: /c:/Go_Tutorial/controllers/accountControllers.go
package controllers

import (
	"log"
	"net/http"
	"webapp/initializer"
	models "webapp/models/db_models"
	"github.com/gin-gonic/gin"
)

	

// CreateStudent godoc
// @Summary Create a new student
// @Description Create a new student with the input payload
// @Success 200 {object} model.Student
// @Router /students [post]
func CreateStudent(c *gin.Context) {

	account := c.ShouldBind(&models.Account{})

	result := initializer.DB.Create(&account)

	if condition := result.Error; condition != nil {
		log.Fatal(condition)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when creating student"})
	}

	c.JSON(http.StatusOK, account)

}

// GetStudents godoc
// @Summary Get all students
// @Description Get all students
// @Tags students
// @Produce  json
// @Success 200 {array} model.Student
// @Router /students [get]
func GetStudents(c *gin.Context) {

	var accounts []models.Account
	initializer.DB.Find(&accounts)
	c.JSON(http.StatusOK, accounts)
}

// DeleteStudent godoc
// @Summary Delete a student
// @Description Delete a student
// @Tags students
// @Param id path int true "Student ID"
// @Success 200 {string} string "Student deleted"
// @Router /students/{id} [delete]
func DeleteStudent(c *gin.Context) {

	result := initializer.DB.Where("id = ?", c.Param("id")).Delete(&models.Account{})
	
	if condition := result.Error; condition != nil {
		log.Fatal(condition)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when deleting student"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}
