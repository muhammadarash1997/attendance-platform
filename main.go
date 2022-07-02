package main

import (
	"attendance-platform/auth"
	"attendance-platform/controllers"
	"attendance-platform/db"
	_ "attendance-platform/docs" // This line is necessary for go-swagger to find your docs!
	"attendance-platform/repositories"
	"attendance-platform/services"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	
	db := db.StartConnection()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	authService := auth.NewService()
	userHandler := controllers.NewUserHandler(userService, authService)

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

	// For User
	router.POST("/api/user/register", userHandler.RegisterUserHandler)                                                        // Register
	router.POST("/api/user/login", userHandler.LoginHandler)                                                                  // Login
	router.GET("/api/user/logout", userHandler.AuthenticateHandler, userHandler.LogoutHandler)                                // Logout
	router.GET("/api/attendance/checkin", userHandler.AuthenticateHandler, attendanceHandler.CheckInHandler)                  // Check In
	router.GET("/api/attendance/checkout/:attendance_id", userHandler.AuthenticateHandler, attendanceHandler.CheckOutHandler) // Check Out
	router.POST("/api/activity", userHandler.AuthenticateHandler, activityHandler.CreateActivityHandler)                      // Menambah aktivitas
	router.PUT("/api/activity", userHandler.AuthenticateHandler, activityHandler.UpdateActivityHandler)                       // Mengedit aktivitas
	router.DELETE("/api/activity/:activity_id", userHandler.AuthenticateHandler, activityHandler.DeleteActivityHandler)       // Menghapus aktivitas
	router.GET("/api/attendance", userHandler.AuthenticateHandler, attendanceHandler.GetAllUserAttendanceHandler)             // Melihat riwayat absensi
	router.GET("/api/activity/:date", userHandler.AuthenticateHandler, activityHandler.GetUserActivitiesByDateHandler)        // Melihat riwayat aktivitas berdasarkan tanggal

	router.Run()
}
