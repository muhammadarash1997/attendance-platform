package controllers

import (
	"attendance-platform/auth"
	"attendance-platform/domain"
	"attendance-platform/dto"
	"attendance-platform/services"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	RegisterEmployeeHandler (*gin.Context)
	LoginHandler        (*gin.Context)
	LogoutHandler       (*gin.Context)
	AuthenticateHandler (*gin.Context)
}

type employeeHandler struct {
	employeeService services.EmployeeService
	authService auth.Service
}

func NewEmployeeHandler(employeeService services.EmployeeService, authService auth.Service) *employeeHandler {
	return &employeeHandler{employeeService, authService}
}

// swagger:route POST /api/employee/register employee registerEmployee
// Create employee
//
// responses:
//		200: registerEmployee
//		422: errorResponse
//		500: errorResponse

func (this *employeeHandler) RegisterEmployeeHandler(c *gin.Context) {
	var registerRequest dto.RegisterRequest

	err := c.ShouldBindJSON(&registerRequest)
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

	err = this.employeeService.RegisterEmployee(registerRequest)
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
		Data:   "Registration success",
	})
}

// swagger:route POST /api/employee/login employee loginEmployee
// Logs employee into the system
//
// responses:
//		200: loginEmployee
//		422: errorResponse
//		500: errorResponse

func (this *employeeHandler) LoginHandler(c *gin.Context) {
	var loginRequest dto.LoginRequest

	err := c.ShouldBindJSON(&loginRequest)
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

	loginResponse, err := this.employeeService.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	tokenGenerated, err := this.authService.GenerateToken(loginResponse.Employee.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}
	loginResponse.Employee.Token = tokenGenerated

	c.JSON(http.StatusOK, dto.Message{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   loginResponse,
	})
}

// swagger:route GET /api/employee/logout employee logoutEmployee
// Logs out employee from the system
//
// Security:
// - Bearer:
// responses:
//		200: logoutEmployee
//		500: errorResponse

func (this *employeeHandler) LogoutHandler(c *gin.Context) {
	currentEmployee := c.MustGet("currentEmployee").(domain.Employee)

	// Rotating id and generate token
	var reverseID []byte
	for i := len(currentEmployee.ID) - 1; i >= 0; i-- {
		reverseID = append(reverseID, currentEmployee.ID[i])
	}

	tokenGenerated, err := this.authService.GenerateToken(string(reverseID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	var logoutResponse dto.LogoutResponse
	logoutResponse.Employee.Token = tokenGenerated

	c.JSON(http.StatusOK, dto.Message{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   logoutResponse,
	})
}

func (this *employeeHandler) AuthenticateHandler(c *gin.Context) {
	// Ambil token dari header
	tokenInput := c.GetHeader("Authorization")

	// Validasi apakah benar itu adalah bearer token
	if !strings.Contains(tokenInput, "Bearer") {
		err := errors.New("Invalid token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Message{
			Code:   http.StatusUnauthorized,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	tokenWithoutBearer := strings.Split(tokenInput, " ")[1]

	// Validasi token apakah berlaku
	token, err := this.authService.ValidateToken(tokenWithoutBearer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Message{
			Code:   http.StatusUnauthorized,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	// Mengubah token yang bertipe jwt.Token menjadi bertipe jwt.MapClaims
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		err = errors.New("Invalid token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Message{
			Code:   http.StatusUnauthorized,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	id := claim["employee_uuid"].(string)
	employee, err := this.employeeService.GetEmployee(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Message{
			Code:   http.StatusUnauthorized,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	c.Set("currentEmployee", employee)
}
