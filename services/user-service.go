package services

import (
	"attendance-platform/domain"
	"attendance-platform/dto"
	"attendance-platform/repositories"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(dto.RegisterRequest) error
	Login(dto.LoginRequest) (dto.LoginResponse, error)
	GetUser(string) (domain.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

func (this *userService) RegisterUser(registerRequest dto.RegisterRequest) error {
	var user domain.User

	// Mapping RegisterRequest to User
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.MinCost)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Error generating password hash")
		return err
	}
	user.SetUsername(registerRequest.Username)
	user.SetName(registerRequest.Name)
	user.SetPasswordHash(string(passwordHash))

	err = this.repository.Save(user)
	if err != nil {
		return err
	}

	return nil
}

func (this *userService) Login(loginRequest dto.LoginRequest) (dto.LoginResponse, error) {
	var loginResponse dto.LoginResponse
	var userDTO dto.UserDTO

	username := loginRequest.Username
	password := loginRequest.Password

	user, err := this.repository.FindByUsername(username)
	if err != nil {
		return loginResponse, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.GetPasswordHash()), []byte(password))
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Wrong password")
		return loginResponse, err
	}

	// Mapping User to UserDTO
	userDTO.ID = user.GetID()
	userDTO.Username = user.GetUsername()
	userDTO.Name = user.GetName()

	// Wrapping UserDTO to LoginResponse
	loginResponse.User = userDTO

	return loginResponse, nil
}

func (this *userService) GetUser(userID string) (domain.User, error) {
	user, err := this.repository.GetByID(userID)
	if err != nil {
		return user, err
	}

	return user, nil
}
