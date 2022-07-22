package repositories

import (
	"github.com/muhammadarash1997/attendance-platform/domain"
	"errors"
	"log"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	Save(domain.Activity) (domain.Activity, error)
	UpdateActivity(domain.Activity) (domain.Activity, error)
	Delete(string) error
	GetByID(string) (domain.Activity, error)
	GetActivitiesByAttendanceID(attendanceID string) ([]domain.Activity, error)
}

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *activityRepository {
	return &activityRepository{db}
}

func (this *activityRepository) Save(activity domain.Activity) (domain.Activity, error) {
	err := this.db.Create(&activity).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Could not create activity")
		return activity, err
	}

	return activity, nil
}

func (this *activityRepository) UpdateActivity(activity domain.Activity) (domain.Activity, error) {
	// Cek apakah data yang ingin di update ada
	_, err := this.GetByID(activity.ID)
	if err != nil {
		return activity,err
	}

	// Lakukan update
	err = this.db.Model(&activity).Updates(activity).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Could not update activity")
		return activity, err
	}

	return activity, nil
}

func (this *activityRepository) Delete(activityID string) error {
	// Cek apakah data yang ingin di delete ada
	_, err := this.GetByID(activityID)
	if err != nil {
		return err
	}

	// Lakukan delete
	err = this.db.Where("id = ?", activityID).Delete(&domain.Activity{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (this *activityRepository) GetByID(id string) (domain.Activity, error) {
	var activity domain.Activity
	err := this.db.Where("id = ?", id).First(&activity).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Activity not found")
		return activity, err
	}

	return activity, nil
}

func (this *activityRepository) GetActivitiesByAttendanceID(attendanceID string) ([]domain.Activity, error) {
	var activities []domain.Activity

	err := this.db.Where("attendance_id = ?", attendanceID).Find(&activities).Error
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Activities not found")
		return activities, err
	}

	return activities, nil
}