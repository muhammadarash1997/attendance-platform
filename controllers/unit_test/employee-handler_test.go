package unit_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	auth "github.com/muhammadarash1997/attendance-platform/auth/mock"
	"github.com/muhammadarash1997/attendance-platform/controllers"
	"github.com/muhammadarash1997/attendance-platform/dto"
	services "github.com/muhammadarash1997/attendance-platform/services/mock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterEmployeeHandler(t *testing.T) {
	t.Run("Register Employee Handler Success", func(t *testing.T) {
		authServiceMock := auth.NewAuthServiceMock()
		employeeServiceMock := services.NewEmployeeServiceMock()
		employeeHandler := controllers.NewEmployeeHandler(employeeServiceMock, authServiceMock)

		registerRequestMock := dto.RegisterRequest{
			Username: "arash",
			Name:     "Muhammad Arash",
			Password: "arashpassword",
		}
		employeeServiceMock.On("RegisterEmployee", registerRequestMock).Return(nil)

		var requestBody bytes.Buffer
		err := json.NewEncoder(&requestBody).Encode(registerRequestMock)
		if err != nil {
			t.Fatal(err)
		}

		router := gin.Default()
		router.POST("/api/employee/register", employeeHandler.RegisterEmployeeHandler)

		request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/employee/register", &requestBody)
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Add("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		response := recorder.Result()

		assert.Equal(t, http.StatusOK, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, float64(http.StatusOK), responseBody["code"])
		assert.Equal(t, "Registration success", responseBody["data"])
		assert.Equal(t, "Ok", responseBody["status"])
	})

	t.Run("Register Employee Handler Failed - Invalid payload", func(t *testing.T) {
		authServiceMock := auth.NewAuthServiceMock()
		employeeServiceMock := services.NewEmployeeServiceMock()
		employeeHandler := controllers.NewEmployeeHandler(employeeServiceMock, authServiceMock)

		var requestBody bytes.Buffer
		err := json.NewEncoder(&requestBody).Encode("")
		if err != nil {
			t.Fatal(err)
		}

		router := gin.Default()
		router.POST("/api/employee/register", employeeHandler.RegisterEmployeeHandler)

		request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/employee/register", &requestBody)
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Add("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		response := recorder.Result()

		assert.Equal(t, http.StatusUnprocessableEntity, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, float64(http.StatusUnprocessableEntity), responseBody["code"])
		assert.Equal(t, "Invalid payload", responseBody["data"])
		assert.Equal(t, "Error", responseBody["status"])
	})

	t.Run("Register Employee Handler Failed - Could not register employee", func(t *testing.T) {
		authServiceMock := auth.NewAuthServiceMock()
		employeeServiceMock := services.NewEmployeeServiceMock()
		employeeHandler := controllers.NewEmployeeHandler(employeeServiceMock, authServiceMock)

		registerRequestMock := dto.RegisterRequest{
			Username: "arash",
			Name:     "Muhammad Arash",
			Password: "arashpassword",
		}
		employeeServiceMock.On("RegisterEmployee", registerRequestMock).Return(errors.New("Error registering employee"))

		var requestBody bytes.Buffer
		err := json.NewEncoder(&requestBody).Encode(registerRequestMock)
		if err != nil {
			t.Fatal(err)
		}

		router := gin.Default()
		router.POST("/api/employee/register", employeeHandler.RegisterEmployeeHandler)

		request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/employee/register", &requestBody)
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Add("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)
		response := recorder.Result()

		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, float64(http.StatusInternalServerError), responseBody["code"])
		assert.Equal(t, "Error registering employee", responseBody["data"])
		assert.Equal(t, "Error", responseBody["status"])
	})
}
