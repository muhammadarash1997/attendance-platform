package mock

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/mock"
)

type authServiceMock struct {
	mock.Mock
}

func NewAuthServiceMock() *authServiceMock {
	return &authServiceMock{}
}

func (this *authServiceMock) GenerateToken(userUUID string) (string, error) {
	args := this.Called(userUUID)

	var tokenGenerated string
	if args.Get(0) != nil {
		tokenGenerated = args.Get(0).(string)
	}

	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return tokenGenerated, err
}

func (this *authServiceMock) ValidateToken(encodedToken string) (*jwt.Token, error) {
	args := this.Called(encodedToken)

	var token *jwt.Token
	if args.Get(0) != nil {
		token = args.Get(0).(*jwt.Token)
	}

	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return token, err
}
