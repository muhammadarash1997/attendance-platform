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
	GetUserAttendanceByDate(string, string) ([]domain.Attendance, error)
	GetAllUserAttendance(string) ([]domain.Attendance, error)
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

func (this *attendanceRepository) GetLatest(userID string) (domain.Attendance, error) {
	var attendance domain.Attendance
	err := this.db.Where("user_id = ?", userID).Order("in_date DESC").First(&attendance).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Could not find latest attendance")
		return attendance, err
	}

	return attendance, nil
}

func (this *attendanceRepository) GetUserAttendanceByDate(userID string, date string) ([]domain.Attendance, error) {
	attendances := []domain.Attendance{}

	lower := fmt.Sprintf("%v 00:00:00", date)
	upper := fmt.Sprintf("%v 24:00:00", date)

	err := this.db.Raw("SELECT * FROM attendances WHERE user_id = ? AND in_date BETWEEN ? AND ? ORDER BY in_date DESC", userID, lower, upper).Scan(&attendances).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Attendance not found")
		return attendances, err
	}

	return attendances, nil
}

// func (this *attendanceRepository) GetUserAttendanceByDate(userID string, date string) (domain.Attendance, error) {
// 	attendance := domain.Attendance{}

// 	lower := fmt.Sprintf("%v 00:00:00", date)
// 	upper := fmt.Sprintf("%v 24:00:00", date)

// 	err := this.db.Raw("SELECT * FROM attendances WHERE user_id = ? AND in_date BETWEEN ? AND ? ORDER BY in_date DESC LIMIT 1", userID, lower, upper).Scan(&attendance).Error
// 	if err != nil {
// 		log.Printf("Error %v", err)
// 		err = errors.New("Attendance not found")
// 		return attendance, err
// 	}

// 	return attendance, nil
// }

func (this *attendanceRepository) GetAllUserAttendance(userID string) ([]domain.Attendance, error) {
	attendances := []domain.Attendance{}

	err := this.db.Where("user_id = ?", userID).Find(&attendances).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Attendances not found")
		return attendances, err
	}

	return attendances, nil
}