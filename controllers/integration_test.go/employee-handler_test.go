package integration_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// 	"github.com/muhammadarash1997/attendance-platform/auth"
// 	"github.com/muhammadarash1997/attendance-platform/controllers"
// 	"github.com/muhammadarash1997/attendance-platform/domain"
// 	"github.com/muhammadarash1997/attendance-platform/dto"
// 	"github.com/muhammadarash1997/attendance-platform/repositories"
// 	"github.com/muhammadarash1997/attendance-platform/services"
// 	"github.com/muhammadarash1997/attendance-platform/utility"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func StartConnectionTest() *gorm.DB {
// 	godotenv.Load("../../.env")
// 	dbHost := os.Getenv("DB_TEST_HOST")
// 	dbPort := os.Getenv("DB_TEST_PORT")
// 	dbUser := os.Getenv("DB_TEST_USER")
// 	dbPass := os.Getenv("DB_TEST_PASS")
// 	dbName := os.Getenv("DB_TEST_NAME")

// 	// jika menggunakan heroku maka sslmode harus require (sslmode=require), jika tidak maka sslmode=disable
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Println(err)
// 		fmt.Println("Failed to connect to test database")
// 		return nil
// 	}
// 	fmt.Println("Success to connect to test database")

// 	db.AutoMigrate(&domain.Employee{})
// 	return db
// }

// func truncateEmployee(db *gorm.DB) {
// 	db.Exec("TRUNCATE employees")
// }

// func TestRegisterEmployeeHandler(t *testing.T) {
// 	t.Run("Register Employee Handler Success", func(t *testing.T) {
// 		db := StartConnectionTest()
// 		truncateEmployee(db)
// 		employeeRepository := repositories.NewEmployeeRepository(db)
// 		hasher := utility.NewHasher()
// 		employeeService := services.NewEmployeeService(hasher, employeeRepository)
// 		authService := auth.NewAuthService()
// 		employeeHandler := controllers.NewEmployeeHandler(employeeService, authService)

// 		registerRequestMock := dto.RegisterRequest{
// 			Username: "arash",
// 			Name:     "Muhammad Arash",
// 			Password: "arashpassword",
// 		}

// 		var requestBody bytes.Buffer
// 		err := json.NewEncoder(&requestBody).Encode(registerRequestMock)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		router := gin.Default()
// 		router.POST("/api/employee/register", employeeHandler.RegisterEmployeeHandler)
// 		request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/employee/register", &requestBody)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		request.Header.Add("Content-Type", "application/json")
// 		recorder := httptest.NewRecorder()
// 		router.ServeHTTP(recorder, request)
// 		response := recorder.Result()

// 		assert.Equal(t, http.StatusOK, response.StatusCode)

// 		body, _ := io.ReadAll(response.Body)
// 		var responseBody map[string]interface{}
// 		json.Unmarshal(body, &responseBody)

// 		assert.Equal(t, float64(http.StatusOK), responseBody["code"])
// 		assert.Equal(t, "Registration success", responseBody["data"])
// 		assert.Equal(t, "Ok", responseBody["status"])
// 	})

// 	t.Run("Register Employee Handler Failed - Username has been used", func(t *testing.T) {
// 		db := StartConnectionTest()
// 		truncateEmployee(db)
// 		employeeRepository := repositories.NewEmployeeRepository(db)
// 		hasher := utility.NewHasher()
// 		employeeService := services.NewEmployeeService(hasher, employeeRepository)
// 		authService := auth.NewAuthService()
// 		employeeHandler := controllers.NewEmployeeHandler(employeeService, authService)

// 		registerRequestMock := dto.RegisterRequest{
// 			Username: "arash",
// 			Name:     "Muhammad Arash",
// 			Password: "arashpassword",
// 		}

// 		var recorder *httptest.ResponseRecorder
// 		var response *http.Response
// 		for i := 0; i < 2; i++{
// 			var requestBody bytes.Buffer
// 			err := json.NewEncoder(&requestBody).Encode(registerRequestMock)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			router := gin.Default()
// 			router.POST("/api/employee/register", employeeHandler.RegisterEmployeeHandler)
// 			request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/employee/register", &requestBody)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			request.Header.Add("Content-Type", "application/json")
// 			recorder = httptest.NewRecorder()
// 			router.ServeHTTP(recorder, request)
// 		}
// 		response = recorder.Result()

// 		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

// 		body, _ := io.ReadAll(response.Body)
// 		var responseBody map[string]interface{}
// 		json.Unmarshal(body, &responseBody)

// 		assert.Equal(t, float64(http.StatusInternalServerError), responseBody["code"])
// 		assert.Equal(t, "Username has been used", responseBody["data"])
// 		assert.Equal(t, "Error", responseBody["status"])
// 	})
// }
