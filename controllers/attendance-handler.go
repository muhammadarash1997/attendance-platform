package controllers

import (
	"net/http"

	"github.com/muhammadarash1997/attendance-platform/domain"
	"github.com/muhammadarash1997/attendance-platform/dto"
	"github.com/muhammadarash1997/attendance-platform/services"

	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	attendanceService services.AttendanceService
}

func NewAttendanceHandler(attendanceService services.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{attendanceService}
}

// swagger:route GET /api/attendance/checkin attendance checkIn
// Check in for attendance
//
// Security:
// - Bearer:
// responses:
//		200: checkIn
//		500: errorResponse

func (this *AttendanceHandler) CheckInHandler(c *gin.Context) {
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

func (this *AttendanceHandler) CheckOutHandler(c *gin.Context) {
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

func (this *AttendanceHandler) GetAllEmployeeAttendanceHandler(c *gin.Context) {
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
