package dto

type LoginResponse struct {
	Employee EmployeeDTO `json:"employee"`
}

type LogoutResponse struct {
	Employee EmployeeDTO `json:"employee"`
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

type GetEmployeeActivitiesByDateResponse struct {
	Activities []ActivityDTO `json:"activities"`
}

type GetAllEmployeeAttendanceResponse struct {
	Attendances []AttendanceDTO `json:"attendances"`
}