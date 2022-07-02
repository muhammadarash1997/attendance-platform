package services

import (
	"attendance-platform/domain"
	"attendance-platform/dto"
	"attendance-platform/repositories"
	"errors"
	"time"
)

type AttendanceService interface {
	CheckIn(string) (dto.CheckInResponse, error)
	CheckOut(string) (dto.CheckOutResponse, error)
	GetAllUserAttendance(string) (dto.GetAllUserAttendanceResponse, error)
}

type attendanceService struct {
	attendanceRepository repositories.AttendanceRepository
}

func NewAttendanceService(attendanceRepository repositories.AttendanceRepository) *attendanceService {
	return &attendanceService{attendanceRepository}
}

func (this *attendanceService) CheckIn(userID string) (dto.CheckInResponse, error) {
	var checkInResponse dto.CheckInResponse
	var attendanceDTO dto.AttendanceDTO
	var attendance domain.Attendance

	now := time.Now()
	attendance.SetUserID(userID)
	attendance.SetInDate(&now)

	// Check have checked out or not
	attendanceLatest, err := this.attendanceRepository.GetLatest(userID)
	if err == nil {
		if attendanceLatest.GetOutDate() == nil {
			err = errors.New("Must check out first")
			return checkInResponse, err
		}
	}

	attendanceAdded, err := this.attendanceRepository.Save(attendance)
	if err != nil {
		return checkInResponse, err
	}

	// Mapping Attendance to AttendanceDTO
	attendanceDTO.ID = attendanceAdded.GetID()
	attendanceDTO.UserID = attendanceAdded.GetUserID()
	attendanceDTO.InDate = attendanceAdded.GetInDate()
	attendanceDTO.OutDate = attendanceAdded.GetOutDate()

	// Mapping AttendanceDTO to CheckInResponse
	checkInResponse.Attendance = attendanceDTO

	return checkInResponse, nil
}

func (this *attendanceService) CheckOut(attendanceID string) (dto.CheckOutResponse, error) {
	var checkOutResponse dto.CheckOutResponse
	var attendanceDTO dto.AttendanceDTO

	// Check logout or not
	attendance, err := this.attendanceRepository.GetByID(attendanceID)
	if err != nil {
		return checkOutResponse, err
	}
	if attendance.GetOutDate() != nil {
		err = errors.New("You have been checked out")
		return checkOutResponse, err
	}

	// Set date of check out
	now := time.Now()
	attendance.SetOutDate(&now)

	// Update attendance
	attendanceUpdated, err := this.attendanceRepository.UpdateAttendance(attendance)
	if err != nil {
		return checkOutResponse, err
	}

	// Mapping Attendance to AttendanceDTO
	attendanceDTO.ID = attendanceUpdated.GetID()
	attendanceDTO.UserID = attendanceUpdated.GetUserID()
	attendanceDTO.InDate = attendanceUpdated.GetInDate()
	attendanceDTO.OutDate = attendanceUpdated.GetOutDate()

	// Wrapping AttendanceDTO to CheckOutResponse
	checkOutResponse.Attendance = attendanceDTO

	return checkOutResponse, err
}

func (this *attendanceService) GetAllUserAttendance(userID string) (dto.GetAllUserAttendanceResponse, error) {
	var getAllUserAttendanceResponse dto.GetAllUserAttendanceResponse
	var attendancesDTO []dto.AttendanceDTO

	attendances, err := this.attendanceRepository.GetAllUserAttendance(userID)
	if err != nil {
		return getAllUserAttendanceResponse, err
	}

	// Mapping []Attendance to []AttendanceDTO
	for _, d := range attendances {
		attendanceDTO := dto.AttendanceDTO{
			ID: d.ID,
			UserID: d.UserID,
			InDate: d.InDate,
			OutDate: d.OutDate,
		}
		attendancesDTO = append(attendancesDTO, attendanceDTO)
	}

	// Wrapping AttendanceDTO to GetAllUserAttendanceResponse
	getAllUserAttendanceResponse.Attendances = attendancesDTO

	return getAllUserAttendanceResponse, nil
}
