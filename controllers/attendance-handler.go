package controllers

import (
	"attendance-platform/domain"
	"attendance-platform/dto"
	"attendance-platform/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AttendanceHandler interface {
	CheckInHandler(*gin.Context)
	CheckOutHandler(*gin.Context)
	GetAllUserAttendanceHandler(*gin.Context)
}

type attendanceHandler struct {
	attendanceService services.AttendanceService
}

func NewAttendanceHandler(attendanceService services.AttendanceService) *attendanceHandler {
	return &attendanceHandler{attendanceService}
}

// swagger:route GET /api/attendance/checkin attendance checkIn
// Check in for attendance
//
// Security:
// - Bearer:
// responses:
//		200: checkIn
//		500: errorResponse

func (this *attendanceHandler) CheckInHandler(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(domain.User)

	checkInResponse, err := this.attendanceService.CheckIn(currentUser.GetID())
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
		Data:   checkInResponse,
	})
}

// swagger:route GET /api/attendance/checkout/{attendance_id} attendance checkOut
// Check out for attendance
//
// Security:
// - Bearer:
// responses:
//		200: checkOut
//		500: errorResponse

func (this *attendanceHandler) CheckOutHandler(c *gin.Context) {
	attendanceID := c.Params.ByName("attendance_id")

	checkInResponse, err := this.attendanceService.CheckOut(attendanceID)
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
		Data:   checkInResponse,
	})
}

// swagger:route GET /api/attendance attendance getAllUserAttendance
// Get all attendance of user
//
// Security:
// - Bearer:
// responses:
//		200: getAllUserAttendance
//		500: errorResponse

func (this *attendanceHandler) GetAllUserAttendanceHandler(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(domain.User)

	getAllUserAttendanceResponse, err := this.attendanceService.GetAllUserAttendance(currentUser.ID)
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
		Data:   getAllUserAttendanceResponse,
	})
}