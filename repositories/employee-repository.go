package repositories

import (
	"github.com/muhammadarash1997/attendance-platform/domain"
	"errors"
	"log"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Save(domain.Employee) error
	FindByUsername(string) (domain.Employee, error)
	GetByID(string) (domain.Employee, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *employeeRepository {
	return &employeeRepository{db}
}

func (this *employeeRepository) Save(employee domain.Employee) error {
	_, err := this.FindByUsername(employee.GetUsername())
	if err == nil {
		return errors.New("Username has been used")
	}

	err = this.db.Create(&employee).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Could not create employee")
		return err
	}

	return nil
}

func (this *employeeRepository) FindByUsername(username string) (domain.Employee, error) {
	employee := domain.Employee{}

	err := this.db.Where("username = ?", username).First(&employee).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Username has not been registered")
		return employee, err
	}

	return employee, nil
}

func (this *employeeRepository) GetByID(employeeID string) (domain.Employee, error) {
	employee := domain.Employee{}

	err := this.db.First(&employee, "id = ?", employeeID).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Employee not found")
		return employee, err
	}

	return employee, nil
}
