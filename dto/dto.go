package dto

import "time"

type EmployeeDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}

type AttendanceDTO struct {
	ID      string    `json:"id"`
	EmployeeID  string    `json:"employee_id"`
	InDate  *time.Time `json:"in_date"`
	OutDate *time.Time `json:"out_date"`
}

type ActivityDTO struct {
	ID           string `json:"id"`
	EmployeeID       string `json:"employee_id"`
	AttendanceID string `json:"attendance_id"`
	Note         string `json:"note"`
}