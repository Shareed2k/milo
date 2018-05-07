package models

import (
	"github.com/milo/util"
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
)

const (
	UserRoleAdmin = "admin"
	UserRoleUser  = "user"
)

type User struct {
	gorm.Model
	Email             string `json:"email" form:"name" validate:"required,email"`
	Username          string `json:"username" form:"username" validate:"required,max=24,alphanum" gorm:"not null;unique_index"`
	Password          string `json:"password,omitempty" form:"password" validate:"required,min=6,max=32" gorm:"-"`
	Role              string `json:"role" form:"role" validate:"required" gorm:"not null"`
	EncryptedPassword []byte `json:"-" gorm:"not null"`
	APIToken          string `json:"api_token" form:"api_token" gorm:"not null;index"`
}

type UserList struct {
	Items []*User `json:"items"`
}

func (m *User) BeforeCreate() error {
	m.GenerateAPIToken()
	return nil
}

func (m *User) BeforeSave() error {
	if m.Password == "" {
		return nil
	}
	return m.encryptPassword()
}

func (m *User) GenerateAPIToken() {
	m.APIToken = util.RandomString(32)
}

func (m *User) encryptPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	m.EncryptedPassword = hashedPassword
	m.Password = "" // just for extra-good measure

	return nil
}
