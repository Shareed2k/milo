package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type DataCenter struct {
	gorm.Model
	Uuid        string   `json:"uuid" gorm:"not null;unique_index"`
	Name        string   `json:"name" form:"name" validate:"required"`
	Provider    string   `json:"provider" form:"name" validate:"required"`
	Description string   `json:"description" form:"description"`
	RegionID    uint     `json:"region_id"`
	Servers     []Server `json:"servers" gorm:"foreignkey:DataCenterID"`
}

func (m *DataCenter) BeforeCreate() error {
	m.Uuid = fmt.Sprintf("%s", uuid.NewV4())
	return nil
}
