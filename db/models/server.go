package models

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
)

type Server struct {
	gorm.Model
	Uuid        string `json:"uuid" validate:"required" gorm:"not null;unique_index"`
	Description string `json:"description"`
	PrivateIp   string `json:"private_ip" validate:"required" gorm:"not null;index"`
	PublicIp    string `json:"public_ip" validate:"required" gorm:"not null;unique_index"`
}

type ServerList struct {
	Items []*Server `json:"items"`
}

func (m *Server) BeforeCreate() error {
	m.Uuid = fmt.Sprintf("%s", uuid.NewV4())
	return nil
}
