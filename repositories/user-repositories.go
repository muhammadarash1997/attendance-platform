package repositories

import (
	"attendance-platform/domain"
	"errors"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(domain.User) error
	FindByUsername(string) (domain.User, error)
	GetByID(string) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (this *userRepository) Save(user domain.User) error {
	_, err := this.FindByUsername(user.GetUsername())
	if err == nil {
		return errors.New("Username has been used")
	}

	err = this.db.Create(&user).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Could not create user")
		return err
	}

	return nil
}

func (this *userRepository) FindByUsername(username string) (domain.User, error) {
	user := domain.User{}

	err := this.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Username has not been registered")
		return user, err
	}

	return user, nil
}

func (this *userRepository) GetByID(userID string) (domain.User, error) {
	user := domain.User{}

	err := this.db.First(&user, "id = ?", userID).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("User not found")
		return user, err
	}

	return user, nil
}
