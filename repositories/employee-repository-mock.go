package repositories

import (
	"github.com/muhammadarash1997/attendance-platform/domain"

	"github.com/stretchr/testify/mock"
)

type employeeRepositoryMock struct {
	mock.Mock
}

func NewEmployeeRepositoryMock() *employeeRepositoryMock {
	return &employeeRepositoryMock{}
}

func (this *employeeRepositoryMock) GetByID(employeeID string) (domain.Employee, error) {
	args := this.Called(employeeID)

	// You have to make sure first that the data is not nil
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

func (this *employeeRepositoryMock) FindByUsername(username string) (domain.Employee, error) {
	args := this.Called(username)

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

func (this *employeeRepositoryMock) Save(employee domain.Employee) error {
	args := this.Called(employee)

	// You have to make sure first that the data is not nil
	var err error
	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}

	return err
}