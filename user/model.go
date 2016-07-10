package user

import (
	"encoding/hex"
	"errors"
	"github.com/codestand/editor/db"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	Id       int32  `json:"id" db:"id" sql:"unique"`
	LoginId  string `json:"login_id" db:"login_id"`
	Password string `json:"-" db:"password"`
}

func AutoMigrate() {
	db.ORM.AutoMigrate(&User{})
}

// TODO: utilize
func EncryptString(s string) string {
	salt := []byte("TODO: change me")
	converted, _ := scrypt.Key([]byte(s), salt, 16384, 8, 1, 32)
	return hex.EncodeToString(converted[:])
}

func Exist(id string) bool {
	u := User{}
	return !db.ORM.Where("login_id = ?", id).First(&u).RecordNotFound()
}

func Save(u *User) {
	u.Password = EncryptString(u.Password)
	db.ORM.Save(&u) // TODO: error handling
}

func Find(loginId interface{}) (u User, err error) {
	if db.ORM.Where("login_id = ?", loginId).First(&u).RecordNotFound() {
		return u, errors.New("not found")
	}
	return u, err
}

func FindWithPassword(id string, password string) (u User, err error) {
	if db.ORM.Where("login_id = ? and password = ?", id, EncryptString(password)).First(&u).RecordNotFound() {
		return u, errors.New("not found")
	}
	return u, err
}

func AllUsers() []User {
	var users []User
	db.ORM.Find(&users)
	return users
}
