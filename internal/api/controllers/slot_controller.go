package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"strconv"
	"webapp/internal/models/request_models"
	"webapp/internal/services"
)

type SlotResponse struct {
	ID        int              `json:"id"`
	StartTime pgtype.Timestamp `json:"start_time"`
	EndTime   pgtype.Timestamp `json:"end_time"`
	Price     float64          `json:"price"`
	RoomId    int              `json:"room_id"`
	MovieId   int              `json:"movie_id"`
}

type SlotController struct {
	SlotService services.SlotServiceInterface
}

func NewSlotController(slotService services.SlotServiceInterface) *SlotController {
	return &SlotController{
		SlotService: slotService,
	}
}

func (sc *SlotController) CreateSlotHandler(c *gin.Context) {

	var req request_models.CreateSlotRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responseError(err.Error()))
		return
	}

	if err := validate.Struct(&req); err != nil {
		c.JSON(http.StatusOK, responseError(err.Error()))
		return
	}

	createdSlot, err := sc.SlotService.CreateSlot(req)
	if err != nil {
		c.JSON(http.StatusOK, responseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responseSuccess("Slot created successfully", []interface{}{SlotResponse{
		ID:        int(createdSlot.ID),
		StartTime: createdSlot.StartTime,
		EndTime:   createdSlot.EndTime,
		Price:     createdSlot.Price,
		RoomId:    int(createdSlot.RoomID),
		MovieId:   int(createdSlot.MovieID),
	},
	}))

}

func (sc *SlotController) GetAllSlotsByMovieIdHandler(c *gin.Context) {

	movieId, _ := strconv.Atoi(c.Param("movieId"))
	slots, err := sc.SlotService.FindAllSlotByMovieID(movieId, 1, 10)
	if err != nil {
		c.JSON(http.StatusOK, responseError(err.Error()))
	}

	var slotResponses []SlotResponse
	for _, slot := range slots {
		slotResponses = append(slotResponses, SlotResponse{
			ID:        int(slot.ID),
			StartTime: slot.StartTime,
			EndTime:   slot.EndTime,
			Price:     slot.Price,
			RoomId:    int(slot.RoomID),
			MovieId:   int(slot.MovieID),
		})
	}

	c.JSON(http.StatusOK, responseSuccess("Slots fetched successfully", []interface{}{slotResponses}))

}
