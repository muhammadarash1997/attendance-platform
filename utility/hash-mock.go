package utility

import "github.com/stretchr/testify/mock"

type hasherMock struct {
	mock.Mock
}

func NewHasherMock() *hasherMock {
	return &hasherMock{}
}

func (this *hasherMock) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	args := this.Called(password, cost)

	var passwordHash []byte
	if args.Get(0) != nil {
		passwordHash = args.Get(0).([]byte)
	}

	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return passwordHash, err
}

func (this *hasherMock) CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	args := this.Called(hashedPassword, password)

	var err error
	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}

	return err
}