package unit_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"

	"github.com/muhammadarash1997/attendance-platform/domain"
	"github.com/muhammadarash1997/attendance-platform/dto"
	"github.com/muhammadarash1997/attendance-platform/repositories/mock"
	"github.com/muhammadarash1997/attendance-platform/services"
	"github.com/muhammadarash1997/attendance-platform/utility"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestGetEmployee(t *testing.T) {

	t.Run("Get Employee Success", func(t *testing.T) {
		// Register test mock and create service with passing mock
		hasherMock := utility.NewHasherMock()
		employeeRepositoryMock := mock.NewEmployeeRepositoryMock()
		employeeService := services.NewEmployeeService(hasherMock, employeeRepositoryMock)

		// Dummy data
		password := []byte("arashpassword")
		passwordHash, _ := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
		employeeMock := domain.Employee{
			ID:           uuid.New().String(),
			Username:     "arash",
			Name:         "Muhammad Arash",
			PasswordHash: string(passwordHash),
		}

		// Setup expectations
		employeeRepositoryMock.On("GetByID", employeeMock.ID).Return(employeeMock, nil)

		// Test
		employee, err := employeeService.GetEmployee(employeeMock.ID)

		// Check
		assert.Nil(t, err)
		assert.NotNil(t, employee)
		assert.Equal(t, employeeMock.ID, employee.ID)
		assert.Equal(t, employeeMock.Username, employee.Username)
		assert.Equal(t, employeeMock.Name, employee.Name)
		assert.Equal(t, employeeMock.PasswordHash, employee.PasswordHash)
	})

	t.Run("Get Employee Failed - Employee not found", func(t *testing.T) {
		// Register test mock and create service with passing mock
		hasherMock := utility.NewHasherMock()
		employeeRepositoryMock := mock.NewEmployeeRepositoryMock()
		employeeService := services.NewEmployeeService(hasherMock, employeeRepositoryMock)

		// Dummy data
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("arashpassword"), bcrypt.MinCost)
		employeeMock := domain.Employee{
			ID:           uuid.New().String(),
			Username:     "arash",
			Name:         "Muhammad Arash",
			PasswordHash: string(passwordHash),
		}

		// Setup expectations
		employeeRepositoryMock.On("GetByID", employeeMock.ID).Return(domain.Employee{}, errors.New("Employee not found"))

		// Test
		employee, err := employeeService.GetEmployee(employeeMock.ID)

		// Check
		assert.NotNil(t, err)
		assert.NotNil(t, employee)
		assert.NotEqual(t, employeeMock.ID, employee.ID)
		assert.NotEqual(t, employeeMock.Username, employee.Username)
		assert.NotEqual(t, employeeMock.Name, employee.Name)
		assert.NotEqual(t, employeeMock.PasswordHash, employee.PasswordHash)
	})
}

func TestRegisterEmployee(t *testing.T) {
	t.Run("Register Employee Success", func(t *testing.T) {
		// Register test mock and create service with passing mock
		hasherMock := utility.NewHasherMock()
		employeeRepositoryMock := mock.NewEmployeeRepositoryMock()
		employeeService := services.NewEmployeeService(hasherMock, employeeRepositoryMock)

		// Dummy data
		registerRequestMock := dto.RegisterRequest{
			Username: "arash",
			Name:     "Muhammad Arash",
			Password: "arashpassword",
		}

		// Setup expectations
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("arashpassword"), bcrypt.MinCost)
		employee := domain.Employee{
			Username:     "arash",
			Name:         "Muhammad Arash",
			PasswordHash: string(passwordHash),
		}
		hasherMock.On("GenerateFromPassword", []byte(registerRequestMock.Password), bcrypt.MinCost).Return(passwordHash, nil)
		employeeRepositoryMock.On("Save", employee).Return(nil)

		// Test
		err := employeeService.RegisterEmployee(registerRequestMock)

		// Check
		assert.Nil(t, err)
	})

	t.Run("Register Employee Failed - Could not generate password hash", func(t *testing.T) {
		// Register test mock and create service with passing mock
		hasherMock := utility.NewHasherMock()
		employeeRepositoryMock := mock.NewEmployeeRepositoryMock()
		employeeService := services.NewEmployeeService(hasherMock, employeeRepositoryMock)

		// Dummy data
		registerRequestMock := dto.RegisterRequest{
			Username: "arash",
			Name:     "Muhammad Arash",
			Password: "arashpassword",
		}

		// Setup expectations
		hasherMock.On("GenerateFromPassword", []byte(registerRequestMock.Password), bcrypt.MinCost).Return(nil, errors.New("Error generating passing hash"))

		// Test
		err := employeeService.RegisterEmployee(registerRequestMock)

		// Check
		assert.NotNil(t, err)
	})

	t.Run("Register Employee Failed - Could not save employee", func(t *testing.T) {
		// Register test mock and create service with passing mock
		hasherMock := utility.NewHasherMock()
		employeeRepositoryMock := mock.NewEmployeeRepositoryMock()
		employeeService := services.NewEmployeeService(hasherMock, employeeRepositoryMock)

		// Dummy data
		registerRequestMock := dto.RegisterRequest{
			Username: "arash",
			Name:     "Muhammad Arash",
			Password: "arashpassword",
		}

		// Setup expectations
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("arashpassword"), bcrypt.MinCost)
		employee := domain.Employee{
			Username:     "arash",
			Name:         "Muhammad Arash",
			PasswordHash: string(passwordHash),
		}
		hasherMock.On("GenerateFromPassword", []byte(registerRequestMock.Password), bcrypt.MinCost).Return(passwordHash, nil)
		employeeRepositoryMock.On("Save", employee).Return(errors.New("Error saving employee"))

		// Test
		err := employeeService.RegisterEmployee(registerRequestMock)

		// Check
		assert.NotNil(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login Success", func(t *testing.T) {
		// Register test mock and create service with passing mock
		hasherMock := utility.NewHasherMock()
		employeeRepositoryMock := mock.NewEmployeeRepositoryMock()
		employeeService := services.NewEmployeeService(hasherMock, employeeRepositoryMock)

		// Dummy data
		loginRequestMock := dto.LoginRequest{
			Username: "arash",
			Password: "arashpassword",
		}

		// Setup expectations
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("arashpassword"), bcrypt.MinCost)
		employee := domain.Employee{
			ID:           uuid.New().String(),
			Username:     loginRequestMock.Username,
			Name:         "Muhammad Arash",
			PasswordHash: string(passwordHash),
		}
		hasherMock.On("CompareHashAndPassword", []byte(employee.PasswordHash), []byte(loginRequestMock.Password)).Return(nil)
		employeeRepositoryMock.On("FindByUsername", loginRequestMock.Username).Return(employee, nil)

		// Test
		loginResponse, err := employeeService.Login(loginRequestMock)

		// Check
		assert.Nil(t, err)
		assert.NotNil(t, loginResponse)
		assert.Equal(t, employee.ID, loginResponse.Employee.ID)
		assert.Equal(t, employee.Username, loginResponse.Employee.Username)
		assert.Equal(t, employee.Name, loginResponse.Employee.Name)
		assert.Equal(t, "", loginResponse.Employee.Token)
	})

	t.Run("Login Failed - Username has not been registered", func(t *testing.T) {
		// Register test mock and create service with passing mock
		hasherMock := utility.NewHasherMock()
		employeeRepositoryMock := mock.NewEmployeeRepositoryMock()
		employeeService := services.NewEmployeeService(hasherMock, employeeRepositoryMock)

		// Dummy data
		loginRequestMock := dto.LoginRequest{
			Username: "arash",
			Password: "arashpassword",
		}
		employee := domain.Employee{}

		// Setup expectations
		employeeRepositoryMock.On("FindByUsername", loginRequestMock.Username).Return(employee, errors.New("Username has not been registered"))

		// Test
		loginResponse, err := employeeService.Login(loginRequestMock)

		// Check
		assert.NotNil(t, err)
		assert.NotNil(t, loginResponse.Employee)
		assert.Equal(t, employee.ID, loginResponse.Employee.ID)
		assert.Equal(t, employee.Username, loginResponse.Employee.Username)
		assert.Equal(t, employee.Name, loginResponse.Employee.Name)
	})

	t.Run("Login Failed - Wrong password", func(t *testing.T) {
		// Register test mock and create service with passing mock
		hasherMock := utility.NewHasherMock()
		employeeRepositoryMock := mock.NewEmployeeRepositoryMock()
		employeeService := services.NewEmployeeService(hasherMock, employeeRepositoryMock)

		// Dummy data
		loginRequestMock := dto.LoginRequest{
			Username: "arash",
			Password: "",
		}

		// Setup expectations
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("arashpassword"), bcrypt.MinCost)
		employee := domain.Employee{
			ID:           uuid.New().String(),
			Username:     loginRequestMock.Username,
			Name:         "Muhammad Arash",
			PasswordHash: string(passwordHash),
		}
		hasherMock.On("CompareHashAndPassword", []byte(employee.PasswordHash), []byte(loginRequestMock.Password)).Return(errors.New("Wrong password"))
		employeeRepositoryMock.On("FindByUsername", loginRequestMock.Username).Return(employee, nil)

		// Test
		loginResponse, err := employeeService.Login(loginRequestMock)

		// // Check
		// assert.NotNil(t, err)
		// assert.NotNil(t, loginResponse.Employee)
		// assert.Equal(t, "", loginResponse.Employee.ID)
		// assert.Equal(t, "", loginResponse.Employee.Username)
		// assert.Equal(t, "", loginResponse.Employee.Name)
		// assert.Equal(t, "", loginResponse.Employee.Token)
		// Check
		assert.NotNil(t, err)
		assert.NotNil(t, loginResponse.Employee)
		assert.NotEqual(t, employee.ID, loginResponse.Employee.ID)
		assert.NotEqual(t, employee.Username, loginResponse.Employee.Username)
		assert.NotEqual(t, employee.Name, loginResponse.Employee.Name)
	})
}
