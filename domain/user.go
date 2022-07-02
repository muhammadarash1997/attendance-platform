package domain

type User struct {
	ID           string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username     string `gorm:"type:varchar(100)"`
	Name         string `gorm:"type:varchar(100)"`
	PasswordHash string `gorm:"type:text"`
}

func (this *User) GetID() string {
	return this.ID
}

func (this *User) SetUsername(username string) {
	this.Username = username
}

func (this *User) GetUsername() string {
	return this.Username
}

func (this *User) SetName(name string) {
	this.Name = name
}

func (this *User) GetName() string {
	return this.Name
}

func (this *User) SetPasswordHash(passwordHash string) {
	this.PasswordHash = passwordHash
}

func (this *User) GetPasswordHash() string {
	return this.PasswordHash
}