package main

import (
	"github.com/muhammadarash1997/attendance-platform/auth"
	"github.com/muhammadarash1997/attendance-platform/controllers"
	"github.com/muhammadarash1997/attendance-platform/db"
	_ "github.com/muhammadarash1997/attendance-platform/docs" // This line is necessary for go-swagger to find your docs!
	"github.com/muhammadarash1997/attendance-platform/repositories"
	"github.com/muhammadarash1997/attendance-platform/services"
	"github.com/muhammadarash1997/attendance-platform/utility"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db := db.StartConnection()

	employeeRepository := repositories.NewEmployeeRepository(db)
	hasher := utility.NewHasher()
	employeeService := services.NewEmployeeService(hasher, employeeRepository)
	authService := auth.NewService()
	employeeHandler := controllers.NewEmployeeHandler(employeeService, authService)

	attendanceRepository := repositories.NewAttendanceRepository(db)
	attendanceService := services.NewAttendanceService(attendanceRepository)
	attendanceHandler := controllers.NewAttendanceHandler(attendanceService)

	activityRepository := repositories.NewActivityRepository(db)
	activityService := services.NewActivityService(activityRepository, attendanceRepository)
	activityHandler := controllers.NewActivityHandler(activityService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/api/docs", "./swagger")

	// swagger:route GET /api/test test testAPI
	// Test API
	//
	// responses:
	//		200: testAPI

	// For API Testing
	router.GET("/api/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	// For Employee
	router.POST("/api/employee/register", employeeHandler.RegisterEmployeeHandler)                                                // Register
	router.POST("/api/employee/login", employeeHandler.LoginHandler)                                                              // Login
	router.GET("/api/employee/logout", employeeHandler.AuthenticateHandler, employeeHandler.LogoutHandler)                        // Logout
	router.GET("/api/attendance/checkin", employeeHandler.AuthenticateHandler, attendanceHandler.CheckInHandler)                  // Check In
	router.GET("/api/attendance/checkout/:attendance_id", employeeHandler.AuthenticateHandler, attendanceHandler.CheckOutHandler) // Check Out
	router.POST("/api/activity", employeeHandler.AuthenticateHandler, activityHandler.CreateActivityHandler)                      // Menambah aktivitas
	router.PUT("/api/activity", employeeHandler.AuthenticateHandler, activityHandler.UpdateActivityHandler)                       // Mengedit aktivitas
	router.DELETE("/api/activity/:activity_id", employeeHandler.AuthenticateHandler, activityHandler.DeleteActivityHandler)       // Menghapus aktivitas
	router.GET("/api/attendance", employeeHandler.AuthenticateHandler, attendanceHandler.GetAllEmployeeAttendanceHandler)         // Melihat riwayat absensi
	router.GET("/api/activity/:date", employeeHandler.AuthenticateHandler, activityHandler.GetEmployeeActivitiesByDateHandler)    // Melihat riwayat aktivitas berdasarkan tanggal

	router.Run()
}
