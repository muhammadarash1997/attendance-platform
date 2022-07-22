package domain

type Activity struct {
	ID           string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	EmployeeID   string `gorm:"type:uuid"`
	AttendanceID string `gorm:"type:uuid"`
	Note         string `gorm:"type:string"`
}

func (this *Activity) SetID(id string) {
	this.ID = id
}

func (this *Activity) GetID() string {
	return this.ID
}

func (this *Activity) SetEmployeeID(employeeID string) {
	this.EmployeeID = employeeID
}

func (this *Activity) GetEmployeeID() string {
	return this.EmployeeID
}

func (this *Activity) SetAttendanceID(attendanceID string) {
	this.AttendanceID = attendanceID
}

func (this *Activity) GetAttendanceID() string {
	return this.AttendanceID
}

func (this *Activity) SetNote(note string) {
	this.Note = note
}

func (this *Activity) GetNote() string {
	return this.Note
}
