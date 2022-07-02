package services

import (
	"attendance-platform/domain"
	"attendance-platform/dto"
	"attendance-platform/repositories"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type EmployeeService interface {
	RegisterEmployee(dto.RegisterRequest) error
	Login(dto.LoginRequest) (dto.LoginResponse, error)
	GetEmployee(string) (domain.Employee, error)
}

type employeeService struct {
	repository repositories.EmployeeRepository
}

func NewEmployeeService(repository repositories.EmployeeRepository) *employeeService {
	return &employeeService{repository}
}

func (this *employeeService) RegisterEmployee(registerRequest dto.RegisterRequest) error {
	var employee domain.Employee

	// Mapping RegisterRequest to Employee
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.MinCost)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Error generating password hash")
		return err
	}
	employee.SetUsername(registerRequest.Username)
	employee.SetName(registerRequest.Name)
	employee.SetPasswordHash(string(passwordHash))

	err = this.repository.Save(employee)
	if err != nil {
		return err
	}

	return nil
}

func (this *employeeService) Login(loginRequest dto.LoginRequest) (dto.LoginResponse, error) {
	var loginResponse dto.LoginResponse
	var employeeDTO dto.EmployeeDTO

	username := loginRequest.Username
	password := loginRequest.Password

	employee, err := this.repository.FindByUsername(username)
	if err != nil {
		return loginResponse, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(employee.GetPasswordHash()), []byte(password))
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Wrong password")
		return loginResponse, err
	}

	// Mapping Employee to EmployeeDTO
	employeeDTO.ID = employee.GetID()
	employeeDTO.Username = employee.GetUsername()
	employeeDTO.Name = employee.GetName()

	// Wrapping EmployeeDTO to LoginResponse
	loginResponse.Employee = employeeDTO

	return loginResponse, nil
}

func (this *employeeService) GetEmployee(employeeID string) (domain.Employee, error) {
	employee, err := this.repository.GetByID(employeeID)
	if err != nil {
		return employee, err
	}

	return employee, nil
}
