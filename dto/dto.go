package dto

import "time"

type EmployeeDTO struct {
	ID       string `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Token    string `json:"token" validate:"required"`
}

type AttendanceDTO struct {
	ID         string     `json:"id" validate:"required"`
	EmployeeID string     `json:"employee_id" validate:"required"`
	InDate     *time.Time `json:"in_date"`
	OutDate    *time.Time `json:"out_date"`
}

type ActivityDTO struct {
	ID           string `json:"id" validate:"required"`
	EmployeeID   string `json:"employee_id" validate:"required"`
	AttendanceID string `json:"attendance_id" validate:"required"`
	Note         string `json:"note" validate:"required"`
}
