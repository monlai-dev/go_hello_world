package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webapp/internal/models/request_models"
	"webapp/internal/services"
)

type TheaterController struct {
	theaterService services.TheaterServiceInterface
}

func NewTheaterController(theaterService services.TheaterServiceInterface) *TheaterController {
	return &TheaterController{
		theaterService: theaterService,
	}
}

func (tc *TheaterController) CreateTheaterHandler(c *gin.Context) {

	var req request_models.TheaterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	createdTheater, err := tc.theaterService.CreateTheater(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responseSuccess("Theater created successfully", []interface{}{createdTheater}))
}

func (tc *TheaterController) GetAllTheatersHandler(c *gin.Context) {

	theaters, err := tc.theaterService.GetAllTheaters()

	if err != nil {
		c.JSON(http.StatusBadRequest, responseError("Error getting theaters"))
		return
	}

	c.JSON(http.StatusOK, responseSuccess("Theaters retrieved successfully", []interface{}{theaters}))
	
}
