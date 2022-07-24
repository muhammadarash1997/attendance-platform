package mock

import (
	"github.com/muhammadarash1997/attendance-platform/domain"
	"github.com/muhammadarash1997/attendance-platform/dto"
	"github.com/stretchr/testify/mock"
)

type employeeServiceMock struct {
	mock.Mock
}

func NewEmployeeServiceMock() *employeeServiceMock {
	return &employeeServiceMock{}
}

func (this *employeeServiceMock) RegisterEmployee(registerRequest dto.RegisterRequest) error {
	args := this.Called(registerRequest)

	var err error
	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}

	return err
}

func (this *employeeServiceMock) Login(loginRequest dto.LoginRequest) (dto.LoginResponse, error) {
	args := this.Called(loginRequest)

	var loginResponse dto.LoginResponse
	if args.Get(0) != nil {
		loginResponse = args.Get(0).(dto.LoginResponse)
	}

	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return loginResponse, err
}

func (this *employeeServiceMock) GetEmployee(employeedID string) (domain.Employee, error) {
	args := this.Called(employeedID)

	var employee domain.Employee
	if args.Get(0) != nil {
		employee = args.Get(0).(domain.Employee)
	}

	var err error
	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return employee, err
}
