package utility

import "golang.org/x/crypto/bcrypt"

type Hasher interface {
	GenerateFromPassword([]byte, int) ([]byte, error)
	CompareHashAndPassword([]byte, []byte) error
}

type hasher struct {
}

func NewHasher() *hasher {
	return &hasher{}
}

func (this *hasher) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func (this *hasher) CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}