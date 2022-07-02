package repositories

import (
	"attendance-platform/domain"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	Save(domain.Attendance) (domain.Attendance, error)
	GetByID(string) (domain.Attendance, error)
	UpdateAttendance(domain.Attendance) (domain.Attendance, error)
	GetLatest(string) (domain.Attendance, error)
	GetEmployeeAttendanceByDate(string, string) ([]domain.Attendance, error)
	GetAllEmployeeAttendance(string) ([]domain.Attendance, error)
}

type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) *attendanceRepository {
	return &attendanceRepository{db}
}

func (this *attendanceRepository) Save(attendance domain.Attendance) (domain.Attendance, error) {
	err := this.db.Create(&attendance).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Could not create attendance")
		return attendance, err
	}

	return attendance, nil
}

func (this *attendanceRepository) GetByID(attendanceID string) (domain.Attendance, error) {
	attendance := domain.Attendance{}

	err := this.db.First(&attendance, "id = ?", attendanceID).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Attendance not found")
		return attendance, err
	}

	return attendance, nil
}

func (this *attendanceRepository) UpdateAttendance(attendance domain.Attendance) (domain.Attendance, error) {
	err := this.db.Model(&attendance).Updates(attendance).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Could not update attendance")
		return attendance, err
	}

	return attendance, nil
}

func (this *attendanceRepository) GetLatest(employeeID string) (domain.Attendance, error) {
	var attendance domain.Attendance
	err := this.db.Where("employee_id = ?", employeeID).Order("in_date DESC").First(&attendance).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Could not find latest attendance")
		return attendance, err
	}

	return attendance, nil
}

func (this *attendanceRepository) GetEmployeeAttendanceByDate(employeeID string, date string) ([]domain.Attendance, error) {
	attendances := []domain.Attendance{}

	lower := fmt.Sprintf("%v 00:00:00", date)
	upper := fmt.Sprintf("%v 24:00:00", date)

	err := this.db.Raw("SELECT * FROM attendances WHERE employee_id = ? AND in_date BETWEEN ? AND ? ORDER BY in_date DESC", employeeID, lower, upper).Scan(&attendances).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Attendance not found")
		return attendances, err
	}

	return attendances, nil
}

// func (this *attendanceRepository) GetEmployeeAttendanceByDate(employeeID string, date string) (domain.Attendance, error) {
// 	attendance := domain.Attendance{}

// 	lower := fmt.Sprintf("%v 00:00:00", date)
// 	upper := fmt.Sprintf("%v 24:00:00", date)

// 	err := this.db.Raw("SELECT * FROM attendances WHERE employee_id = ? AND in_date BETWEEN ? AND ? ORDER BY in_date DESC LIMIT 1", employeeID, lower, upper).Scan(&attendance).Error
// 	if err != nil {
// 		log.Printf("Error %v", err)
// 		err = errors.New("Attendance not found")
// 		return attendance, err
// 	}

// 	return attendance, nil
// }

func (this *attendanceRepository) GetAllEmployeeAttendance(employeeID string) ([]domain.Attendance, error) {
	attendances := []domain.Attendance{}

	err := this.db.Where("employee_id = ?", employeeID).Find(&attendances).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Attendances not found")
		return attendances, err
	}

	return attendances, nil
}