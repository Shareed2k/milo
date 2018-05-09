package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Server struct {
	gorm.Model
	Uuid         string `json:"uuid" gorm:"not null;unique_index"`
	Name         string `json:"name" form:"name" validate:"required"`
	Description  string `json:"description" form:"description"`
	PrivateIp    string `json:"private_ip" form:"private_ip" validate:"required" gorm:"not null;index"`
	PublicIp     string `json:"public_ip" form:"public_ip" validate:"required" gorm:"not null;unique_index"`
	DataCenterID uint   `json:"server_id"`
}

func (m *Server) BeforeCreate() error {
	m.Uuid = fmt.Sprintf("%s", uuid.NewV4())
	return nil
}
