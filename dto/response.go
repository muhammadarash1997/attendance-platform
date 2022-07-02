package dto

type LoginResponse struct {
	User UserDTO `json:"user"`
}

type LogoutResponse struct {
	User UserDTO `json:"user"`
}

type CreateActivityResponse struct {
	Activity ActivityDTO `json:"activity"`
}

type UpdateActivityResponse struct {
	Activity ActivityDTO `json:"activity"`
}

type CheckInResponse struct {
	Attendance AttendanceDTO `json:"attendance"`
}

type CheckOutResponse struct {
	Attendance AttendanceDTO `json:"attendance"`
}

type GetUserActivitiesByDateResponse struct {
	Activities []ActivityDTO `json:"activities"`
}

type GetAllUserAttendanceResponse struct {
	Attendances []AttendanceDTO `json:"attendances"`
}