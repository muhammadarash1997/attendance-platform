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

type UserHandler struct {
	RegisterUserHandler (*gin.Context)
	LoginHandler        (*gin.Context)
	LogoutHandler       (*gin.Context)
	AuthenticateHandler (*gin.Context)
}

type userHandler struct {
	userService services.UserService
	authService auth.Service
}

func NewUserHandler(userService services.UserService, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// swagger:route POST /api/user/register user registerUser
// Create user
//
// responses:
//		200: registerUser
//		422: errorResponse
//		500: errorResponse

func (this *userHandler) RegisterUserHandler(c *gin.Context) {
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

	err = this.userService.RegisterUser(registerRequest)
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

// swagger:route POST /api/user/login user loginUser
// Logs user into the system
//
// responses:
//		200: loginUser
//		422: errorResponse
//		500: errorResponse

func (this *userHandler) LoginHandler(c *gin.Context) {
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

	loginResponse, err := this.userService.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	tokenGenerated, err := this.authService.GenerateToken(loginResponse.User.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}
	loginResponse.User.Token = tokenGenerated

	c.JSON(http.StatusOK, dto.Message{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   loginResponse,
	})
}

// swagger:route GET /api/user/logout user logoutUser
// Logs out user from the system
//
// Security:
// - Bearer:
// responses:
//		200: logoutUser
//		500: errorResponse

func (this *userHandler) LogoutHandler(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(domain.User)

	// Rotating id and generate token
	var reverseID []byte
	for i := len(currentUser.ID) - 1; i >= 0; i-- {
		reverseID = append(reverseID, currentUser.ID[i])
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
	logoutResponse.User.Token = tokenGenerated

	c.JSON(http.StatusOK, dto.Message{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   logoutResponse,
	})
}

func (this *userHandler) AuthenticateHandler(c *gin.Context) {
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

	id := claim["user_uuid"].(string)
	user, err := this.userService.GetUser(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Message{
			Code:   http.StatusUnauthorized,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	c.Set("currentUser", user)
}
