package dto

type RegisterRequest struct {
	Username     string `json:"username"`
	Name         string `json:"name"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username    string `json:"username"`
	Password string `json:"password"`
}

type CreateActivityRequest struct {
	UserID       string `json:"user_id"`
	AttendanceID string `json:"attendance_id"`
	Note         string `json:"note"`
}

type UpdateActivityRequest struct {
	Activity ActivityDTO	`json:"activity"`
}