package internal

import (
	"github.com/milo/db/models"
)

type ServerRepository interface {
	Create(s *models.Server) (uint, error)
}

type serverRepo struct {
	core Core
	*Database
}

func NewServerRepository(c Core) (Repository, error) {
	db := c.GetMaster().GetDatabase()
	return &serverRepo{c, db}, nil
}

func (r *serverRepo) Create(s *models.Server) (uint, error) {
	if err := r.DB.Create(s).Error; err != nil {
		return 0, err
	}

	return s.ID, nil
}
