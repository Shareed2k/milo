package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Provider struct {
	gorm.Model
	Uuid        string        `json:"uuid" gorm:"not null;unique_index"`
	Name        string        `json:"name" form:"name" validate:"required"`
	Description string        `json:"description" form:"description"`
	DataCenters []*DataCenter `json:"datacenters" gorm:"foreignkey:ProviderID;AssociationForeignKey:ID"`
}

func (m *Provider) BeforeCreate() error {
	m.Uuid = fmt.Sprintf("%s", uuid.NewV4())
	return nil
}
