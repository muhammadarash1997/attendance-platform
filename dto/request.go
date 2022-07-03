package dto

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateActivityRequest struct {
	EmployeeID   string `json:"employee_id" validate:"required"`
	AttendanceID string `json:"attendance_id" validate:"required"`
	Note         string `json:"note" validate:"required"`
}

type UpdateActivityRequest struct {
	Activity ActivityDTO `json:"activity"`
}
