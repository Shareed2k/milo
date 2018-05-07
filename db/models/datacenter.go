package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type DataCenter struct {
	gorm.Model
	Uuid        string `json:"uuid" validate:"required" gorm:"not null;unique_index"`
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description"`
}

type DataCenterList struct {
	Items []*DataCenter `json:"items"`
}

func (m *DataCenter) BeforeCreate() error {
	m.Uuid = fmt.Sprintf("%s", uuid.NewV4())
	return nil
}
