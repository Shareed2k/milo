package models

type User struct {
	ID           uint `gorm:"primary_key"`
	Username     string
	Email        string
	PasswordHash string
}