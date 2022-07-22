package controllers

import (
	"github.com/muhammadarash1997/attendance-platform/domain"
	"github.com/muhammadarash1997/attendance-platform/dto"
	"github.com/muhammadarash1997/attendance-platform/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AttendanceHandler interface {
	CheckInHandler(*gin.Context)
	CheckOutHandler(*gin.Context)
	GetAllEmployeeAttendanceHandler(*gin.Context)
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
	currentEmployee := c.MustGet("currentEmployee").(domain.Employee)

	checkInResponse, err := this.attendanceService.CheckIn(currentEmployee.GetID())
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

// swagger:route GET /api/attendance attendance getAllEmployeeAttendance
// Get all attendance of employee
//
// Security:
// - Bearer:
// responses:
//		200: getAllEmployeeAttendance
//		500: errorResponse

func (this *attendanceHandler) GetAllEmployeeAttendanceHandler(c *gin.Context) {
	currentEmployee := c.MustGet("currentEmployee").(domain.Employee)

	getAllEmployeeAttendanceResponse, err := this.attendanceService.GetAllEmployeeAttendance(currentEmployee.ID)
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
		Data:   getAllEmployeeAttendanceResponse,
	})
}