package auth

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userUUID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (this *service) GenerateToken(userUUID string) (string, error) {
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

func (this *service) ValidateToken(encodedToken string) (*jwt.Token, error) {
	// secretKey := []byte(os.Getenv("SECRET_KEY"))
	secretKey := []byte("S3C123TKEY")

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
