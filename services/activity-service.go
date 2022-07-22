package services

import (
	"github.com/muhammadarash1997/attendance-platform/domain"
	"github.com/muhammadarash1997/attendance-platform/dto"
	"github.com/muhammadarash1997/attendance-platform/repositories"
	"errors"
)

type ActivityService interface {
	CreateActivity(dto.CreateActivityRequest) (dto.CreateActivityResponse, error)
	UpdateActivity(dto.UpdateActivityRequest) (dto.UpdateActivityResponse, error)
	DeleteActivity(string) error
	GetEmployeeActivitiesByDate(string,string) (dto.GetEmployeeActivitiesByDateResponse, error)
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
	activity.SetEmployeeID(createActivityRequest.EmployeeID)
	activity.SetAttendanceID(createActivityRequest.AttendanceID)
	activity.SetNote(createActivityRequest.Note)

	activityAdded, err := this.activityRepository.Save(activity)
	if err != nil {
		return createActivityResponse, err
	}

	// Mapping Activity to ActivityDTO
	activityDTO.ID = activityAdded.GetID()
	activityDTO.EmployeeID = activityAdded.GetEmployeeID()
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

	// Mapping UpdateActivityRequest to Activity
	activity.SetID(updateActivityRequest.Activity.ID)
	activity.SetEmployeeID(updateActivityRequest.Activity.EmployeeID)
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
	activityDTO.EmployeeID = activityUpdated.GetEmployeeID()
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

func (this *activityService) GetEmployeeActivitiesByDate(employeeID string, date string) (dto.GetEmployeeActivitiesByDateResponse, error) {
	var getEmployeeActivitiesByDateResponse dto.GetEmployeeActivitiesByDateResponse
	var activitiesDTO []dto.ActivityDTO

	// Find attendance based on employee id dan date
	attendances, err := this.attendanceRepository.GetEmployeeAttendanceByDate(employeeID, date)
	if err != nil {
		return getEmployeeActivitiesByDateResponse, err
	}

	for _, v := range attendances {
		// Find activities based on attendance id
		activities, err := this.activityRepository.GetActivitiesByAttendanceID(v.ID)
		if err != nil {
			return getEmployeeActivitiesByDateResponse, err
		}

		// Mapping []Activity to []ActivityDTO
		for _, d := range activities {
			activityDTO := dto.ActivityDTO{
				ID: d.ID,
				EmployeeID: d.EmployeeID,
				AttendanceID: d.AttendanceID,
				Note: d.Note,
			}
			activitiesDTO = append(activitiesDTO, activityDTO)
		}
	}

	// Wrapping []ActivityDTO to GetEmployeeActivitiesByDateResponse
	getEmployeeActivitiesByDateResponse.Activities = activitiesDTO

	return getEmployeeActivitiesByDateResponse, nil
}