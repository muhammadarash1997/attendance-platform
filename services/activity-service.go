package services

import (
	"attendance-platform/domain"
	"attendance-platform/dto"
	"attendance-platform/repositories"
	"errors"
	"fmt"
)

type ActivityService interface {
	CreateActivity(dto.CreateActivityRequest) (dto.CreateActivityResponse, error)
	UpdateActivity(dto.UpdateActivityRequest) (dto.UpdateActivityResponse, error)
	DeleteActivity(string) error
	GetUserActivitiesByDate(string,string) (dto.GetUserActivitiesByDateResponse, error)
}

type activityService struct {
	activityRepository   repositories.ActivityRepository
	attendanceRepository repositories.AttendanceRepository
}

func NewActivityService(activityRepository repositories.ActivityRepository, attendanceRepository repositories.AttendanceRepository) *activityService {
	return &activityService{activityRepository, attendanceRepository}
}

func (this *activityService) CreateActivity(createActivityRequest dto.CreateActivityRequest) (dto.CreateActivityResponse, error) {
	var activity domain.Activity
	var activityDTO dto.ActivityDTO
	var createActivityResponse dto.CreateActivityResponse

	// Check logout or not
	attendance, err := this.attendanceRepository.GetByID(createActivityRequest.AttendanceID)
	if err != nil {
		return createActivityResponse, err
	}
	if attendance.GetOutDate() != nil {
		err = errors.New("You have been checked out")
		return createActivityResponse, err
	}

	// Mapping CreateActivityRequest to Activity
	activity.SetUserID(createActivityRequest.UserID)
	activity.SetAttendanceID(createActivityRequest.AttendanceID)
	activity.SetNote(createActivityRequest.Note)

	activityAdded, err := this.activityRepository.Save(activity)
	if err != nil {
		return createActivityResponse, err
	}

	// Mapping Activity to ActivityDTO
	activityDTO.ID = activityAdded.GetID()
	activityDTO.UserID = activityAdded.GetUserID()
	activityDTO.AttendanceID = activityAdded.GetAttendanceID()
	activityDTO.Note = activityAdded.GetNote()

	// Wrapping ActivityDTO to CreateActivityResponse
	createActivityResponse.Activity = activityDTO

	return createActivityResponse, nil
}

func (this *activityService) UpdateActivity(updateActivityRequest dto.UpdateActivityRequest) (dto.UpdateActivityResponse, error) {
	var activity domain.Activity
	var activityDTO dto.ActivityDTO
	var updateActivityResponse dto.UpdateActivityResponse

	fmt.Println("INI adalah id", updateActivityRequest.Activity.AttendanceID)

	// Mapping UpdateActivityRequest to Activity
	activity.SetID(updateActivityRequest.Activity.ID)
	activity.SetUserID(updateActivityRequest.Activity.UserID)
	activity.SetAttendanceID(updateActivityRequest.Activity.AttendanceID)
	activity.SetNote(updateActivityRequest.Activity.Note)

	// Check logout or not
	attendance, err := this.attendanceRepository.GetByID(activity.GetAttendanceID())
	if err != nil {
		return updateActivityResponse, err
	}
	if attendance.GetOutDate() != nil {
		err = errors.New("You have been checked out")
		return updateActivityResponse, err
	}

	activityUpdated, err := this.activityRepository.UpdateActivity(activity)
	if err != nil {
		return updateActivityResponse, err
	}

	// Mapping Activity to ActivityDTO
	activityDTO.ID = activityUpdated.GetID()
	activityDTO.UserID = activityUpdated.GetUserID()
	activityDTO.AttendanceID = activityUpdated.GetAttendanceID()
	activityDTO.Note = activityUpdated.GetNote()

	// Wrapping ActivityDTO to CreateActivityResponse
	updateActivityResponse.Activity = activityDTO

	return updateActivityResponse, nil
}

func (this *activityService) DeleteActivity(activityID string) error {
	// Get Activity to get AttendanceID
	activity, err := this.activityRepository.GetByID(activityID)
	if err != nil {
		return err
	}

	// Check logout or not
	attendance, err := this.attendanceRepository.GetByID(activity.AttendanceID)
	if err != nil {
		return err
	}
	if attendance.GetOutDate() != nil {
		err = errors.New("You have been checked out")
		return err
	}

	err = this.activityRepository.Delete(activityID)
	if err != nil {
		return err
	}

	return nil
}

func (this *activityService) GetUserActivitiesByDate(userID string, date string) (dto.GetUserActivitiesByDateResponse, error) {
	var getUserActivitiesByDateResponse dto.GetUserActivitiesByDateResponse
	var activitiesDTO []dto.ActivityDTO

	// Find attendance based on user id dan date
	attendances, err := this.attendanceRepository.GetUserAttendanceByDate(userID, date)
	if err != nil {
		return getUserActivitiesByDateResponse, err
	}

	for _, v := range attendances {
		// Find activities based on attendance id
		activities, err := this.activityRepository.GetActivitiesByAttendanceID(v.ID)
		if err != nil {
			return getUserActivitiesByDateResponse, err
		}

		// Mapping []Activity to []ActivityDTO
		for _, d := range activities {
			activityDTO := dto.ActivityDTO{
				ID: d.ID,
				UserID: d.UserID,
				AttendanceID: d.AttendanceID,
				Note: d.Note,
			}
			activitiesDTO = append(activitiesDTO, activityDTO)
		}
	}

	// Wrapping []ActivityDTO to GetUserActivitiesByDateResponse
	getUserActivitiesByDateResponse.Activities = activitiesDTO

	return getUserActivitiesByDateResponse, nil
}