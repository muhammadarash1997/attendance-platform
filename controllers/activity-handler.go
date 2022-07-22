package controllers

import (
	"github.com/muhammadarash1997/attendance-platform/domain"
	"github.com/muhammadarash1997/attendance-platform/dto"
	"github.com/muhammadarash1997/attendance-platform/services"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActivityHandler interface {
	CreateActivityHandler(*gin.Context)
	UpdateActivityHandler(c *gin.Context)
	DeleteActivityHandler(c *gin.Context)
	GetEmployeeActivitiesByDateHandler(c *gin.Context)
}

type activityHandler struct {
	activityService services.ActivityService
}

func NewActivityHandler(activityService services.ActivityService) *activityHandler {
	return &activityHandler{activityService}
}

// swagger:route POST /api/activity activity createActivity
// Create an activity of attendance
//
// Security:
// - Bearer:
// responses:
//		200: createActivity
//		422: errorResponse
//		500: errorResponse

func (this *activityHandler) CreateActivityHandler(c *gin.Context) {
	var createActivityRequest dto.CreateActivityRequest

	err := c.ShouldBindJSON(&createActivityRequest)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Invalid payload")
		c.JSON(http.StatusUnprocessableEntity, dto.Message{
			Code:   http.StatusUnprocessableEntity,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	createActivityResponse, err := this.activityService.CreateActivity(createActivityRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Message{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   createActivityResponse,
	})
}

// swagger:route PUT /api/activity activity updateActivity
// Update an activity of attendance
//
// Security:
// - Bearer:
// responses:
//		200: updateActivity
//		422: errorResponse
//		500: errorResponse

func (this *activityHandler) UpdateActivityHandler(c *gin.Context) {
	var updateActivityRequest dto.UpdateActivityRequest

	err := c.ShouldBindJSON(&updateActivityRequest)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Invalid payload")
		c.JSON(http.StatusUnprocessableEntity, dto.Message{
			Code:   http.StatusUnprocessableEntity,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	updateActivityResponse, err := this.activityService.UpdateActivity(updateActivityRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Message{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   updateActivityResponse,
	})
}

// swagger:route DELETE /api/activity/{activity_id} activity deleteActivity
// Delete an activity of attendance
//
// Security:
// - Bearer:
// responses:
//		200: deleteActivity
//		500: errorResponse

func (this *activityHandler) DeleteActivityHandler(c *gin.Context) {
	activityID := c.Params.ByName("activity_id")

	err := this.activityService.DeleteActivity(activityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Message{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   "Deleting activity success",
	})
}

// swagger:route GET /api/activity/{date} activity getEmployeeActivitiesByDate
// Get activities of employee at certain date
//
// Security:
// - Bearer:
// responses:
//		200: getEmployeeActivitiesByDate
//		500: errorResponse

func (this *activityHandler) GetEmployeeActivitiesByDateHandler(c *gin.Context) {
	currentEmployee := c.MustGet("currentEmployee").(domain.Employee)
	date := c.Params.ByName("date")

	activities, err := this.activityService.GetEmployeeActivitiesByDate(currentEmployee.ID, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Message{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   activities,
	})
}
