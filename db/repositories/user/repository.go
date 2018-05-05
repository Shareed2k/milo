package user

import (
	"github.com/milo/db/models"
	"github.com/milo/internal"
)

type UserRepository interface {
	Create(user *models.User) uint
}

type repository struct {
	core internal.Core
	*internal.Database
}

func NewUserRepository(c internal.Core) UserRepository {
	db := c.GetOperator().(internal.MasterOperator).GetDatabase()
	return &repository{c, db}
}

func (r *repository) Create(user *models.User) uint {
	r.DB.Create(user)
	return user.ID
}
