package db_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/muhammadarash1997/attendance-platform/domain"
	"github.com/muhammadarash1997/attendance-platform/repositories"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartConnectionTest() *gorm.DB {
	godotenv.Load("../../.env")
	dbHost := os.Getenv("DB_TEST_HOST")
	dbPort := os.Getenv("DB_TEST_PORT")
	dbUser := os.Getenv("DB_TEST_USER")
	dbPass := os.Getenv("DB_TEST_PASS")
	dbName := os.Getenv("DB_TEST_NAME")

	// jika menggunakan heroku maka sslmode harus require (sslmode=require), jika tidak maka sslmode=disable
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
		fmt.Println("Failed to connect to test database")
		return nil
	}
	fmt.Println("Success to connect to test database")

	db.AutoMigrate(&domain.Employee{})
	return db
}

func truncateEmployee(db *gorm.DB) {
	db.Exec("TRUNCATE employees")
}

func TestSave(t *testing.T) {
	t.Run("Save Success", func(t *testing.T) {
		db := StartConnectionTest()
		truncateEmployee(db)
		employeeRepository := repositories.NewEmployeeRepository(db)

		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("arashpassword"), bcrypt.MinCost)
		employeeMock := domain.Employee{
			ID:           uuid.New().String(),
			Username:     "arash",
			Name:         "Muhammad Arash",
			PasswordHash: string(passwordHash),
		}

		err := employeeRepository.Save(employeeMock)
		assert.Nil(t, err)
	})

	t.Run("Save Failed - Username has been used", func(t *testing.T) {
		db := StartConnectionTest()
		truncateEmployee(db)
		employeeRepository := repositories.NewEmployeeRepository(db)

		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("arashpassword"), bcrypt.MinCost)
		employeeMock := domain.Employee{
			ID:           uuid.New().String(),
			Username:     "arash",
			Name:         "Muhammad Arash",
			PasswordHash: string(passwordHash),
		}
		employeeRepository.Save(employeeMock)

		err := employeeRepository.Save(employeeMock)
		assert.NotNil(t, err)
	})
}

func TestFindByUsername(t *testing.T) {
	t.Run("Find By Username Success", func(t *testing.T) {
		db := StartConnectionTest()
		truncateEmployee(db)
		employeeRepository := repositories.NewEmployeeRepository(db)

		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("arashpassword"), bcrypt.MinCost)
		employeeMock := domain.Employee{
			ID:           uuid.New().String(),
			Username:     "arash",
			Name:         "Muhammad Arash",
			PasswordHash: string(passwordHash),
		}
		employeeRepository.Save(employeeMock)

		employee, err := employeeRepository.FindByUsername(employeeMock.GetUsername())
		assert.Nil(t, err)
		assert.NotNil(t, employee)
		assert.Equal(t, employeeMock.ID, employee.ID)
		assert.Equal(t, employeeMock.Username, employee.Username)
		assert.Equal(t, employeeMock.Name, employee.Name)
		assert.Equal(t, employeeMock.PasswordHash, employee.PasswordHash)
	})

	t.Run("Find By Username Failed - Username has not been registered", func(t *testing.T) {
		db := StartConnectionTest()
		truncateEmployee(db)
		employeeRepository := repositories.NewEmployeeRepository(db)

		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("arashpassword"), bcrypt.MinCost)
		employeeMock := domain.Employee{
			ID:           uuid.New().String(),
			Username:     "arash",
			Name:         "Muhammad Arash",
			PasswordHash: string(passwordHash),
		}

		employee, err := employeeRepository.FindByUsername(employeeMock.GetUsername())
		assert.NotNil(t, err)
		assert.NotNil(t, employee)
		assert.NotEqual(t, employeeMock.ID, employee.ID)
		assert.NotEqual(t, employeeMock.Username, employee.Username)
		assert.NotEqual(t, employeeMock.Name, employee.Name)
		assert.NotEqual(t, employeeMock.PasswordHash, employee.PasswordHash)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Get By ID Success", func(t *testing.T) {
		db := StartConnectionTest()
		truncateEmployee(db)
		employeeRepository := repositories.NewEmployeeRepository(db)

		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("arashpassword"), bcrypt.MinCost)
		employeeMock := domain.Employee{
			ID:           uuid.New().String(),
			Username:     "arash",
			Name:         "Muhammad Arash",
			PasswordHash: string(passwordHash),
		}
		employeeRepository.Save(employeeMock)

		employee, err := employeeRepository.GetByID(employeeMock.ID)
		assert.Nil(t, err)
		assert.NotNil(t, employee)
		assert.Equal(t, employeeMock.ID, employee.ID)
		assert.Equal(t, employeeMock.Username, employee.Username)
		assert.Equal(t, employeeMock.Name, employee.Name)
		assert.Equal(t, employeeMock.PasswordHash, employee.PasswordHash)
	})

	t.Run("Get By ID Failed - Employee not found", func(t *testing.T) {
		db := StartConnectionTest()
		truncateEmployee(db)
		employeeRepository := repositories.NewEmployeeRepository(db)

		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("arashpassword"), bcrypt.MinCost)
		employeeMock := domain.Employee{
			ID:           uuid.New().String(),
			Username:     "arash",
			Name:         "Muhammad Arash",
			PasswordHash: string(passwordHash),
		}

		employee, err := employeeRepository.GetByID(uuid.New().String())
		assert.NotNil(t, err)
		assert.NotNil(t, employee)
		assert.NotEqual(t, employeeMock.ID, employee.ID)
		assert.NotEqual(t, employeeMock.Username, employee.Username)
		assert.NotEqual(t, employeeMock.Name, employee.Name)
		assert.NotEqual(t, employeeMock.PasswordHash, employee.PasswordHash)
	})
}
