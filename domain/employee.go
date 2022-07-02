package domain

type Employee struct {
	ID           string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username     string `gorm:"type:varchar(100)"`
	Name         string `gorm:"type:varchar(100)"`
	PasswordHash string `gorm:"type:text"`
}

func (this *Employee) GetID() string {
	return this.ID
}

func (this *Employee) SetUsername(username string) {
	this.Username = username
}

func (this *Employee) GetUsername() string {
	return this.Username
}

func (this *Employee) SetName(name string) {
	this.Name = name
}

func (this *Employee) GetName() string {
	return this.Name
}

func (this *Employee) SetPasswordHash(passwordHash string) {
	this.PasswordHash = passwordHash
}

func (this *Employee) GetPasswordHash() string {
	return this.PasswordHash
}