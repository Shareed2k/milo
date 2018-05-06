package internal

import "github.com/milo/db/models"

type ServerRepository interface {
	Create(s *models.Server) uint
}

type serverRepo struct {
	core Core
	*Database
}

func NewServerRepository(c Core) (Repository, error) {
	db := c.GetMaster().GetDatabase()
	return &serverRepo{c, db}, nil
}

func (r *serverRepo) Create(s *models.Server) uint {
	r.DB.Create(s)

	return s.ID
}
