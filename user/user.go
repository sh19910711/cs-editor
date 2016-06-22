package user

import (
	"encoding/hex"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	ID       int32  `json:"-"`
	LoginID  string `json:"login_id" form:"user.login_id" db:"login_id" binding:"required,max=32"`
	Password string `json:"-" form:"user.password" db:"password" binding:"required,max=128"`
}

// TODO: make private
func (ua *User) EncryptedPassword() string {
	salt := []byte("TODO: change me")
	converted, _ := scrypt.Key([]byte(ua.Password), salt, 16384, 8, 1, 32)
	return hex.EncodeToString(converted[:])
}
