package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Region struct {
	gorm.Model
	Uuid        string       `json:"uuid" gorm:"not null;unique_index"`
	Name        string       `json:"name" form:"name" validate:"required" gorm:"not null;unique_index"`
	Description string       `json:"description" form:"description"`
	DataCenters []DataCenter `json:"data_centers" gorm:"foreignkey:RegionID"`
}

func (m *Region) BeforeCreate() error {
	m.Uuid = fmt.Sprintf("%s", uuid.NewV4())
	return nil
}
