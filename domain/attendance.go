package domain

import "time"

type Attendance struct {
	ID      string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID  string `gorm:"type:uuid"`
	InDate  *time.Time
	OutDate *time.Time
}

func (this *Attendance) GetID() string {
	return this.ID
}

func (this *Attendance) SetUserID(userID string) {
	this.UserID = userID
}

func (this *Attendance) GetUserID() string {
	return this.UserID
}

func (this *Attendance) SetInDate(inDate *time.Time) {
	this.InDate = inDate
}

func (this *Attendance) GetInDate() *time.Time {
	return this.InDate
}

func (this *Attendance) SetOutDate(outDate *time.Time) {
	this.OutDate = outDate
}

func (this *Attendance) GetOutDate() *time.Time {
	return this.OutDate
}
