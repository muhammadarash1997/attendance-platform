package auth

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	GenerateToken(string) (string, error)
	ValidateToken(string) (*jwt.Token, error)
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (this *authService) GenerateToken(userUUID string) (string, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	tokenHourLifespanString := os.Getenv("TOKEN_HOUR_LIFESPAN")

	tokenHourLifespan, err := strconv.Atoi(tokenHourLifespanString)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Error generating token")
		return "", err
	}

	claim := jwt.MapClaims{}
	claim["user_uuid"] = userUUID
	claim["exp"] = time.Now().Add(time.Hour * time.Duration(tokenHourLifespan)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenGenerated, err := token.SignedString(secretKey)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Error signing token")
		return tokenGenerated, err
	}

	return tokenGenerated, nil
}

func (this *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return secretKey, nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
