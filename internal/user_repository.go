package internal

import (
	"fmt"
	"github.com/milo/db/models"
	"github.com/milo/util"
)

type UserRepository interface {
	DetectOrCreateAdmin() error
}

type repository struct {
	core Core
	*Database
}

func NewUserRepository(c Core) (Repository, error) {
	db := c.GetOperator().(MasterOperator).GetDatabase()
	return &repository{c, db}, nil
}

func (r *repository) DetectOrCreateAdmin() error {
	settings := r.core.GetSettings().GetOptions()
	if settings.SupportPassword != "" {
		support := &models.User{
			Username: "support",
			Password: settings.SupportPassword,
			Role:     models.UserRoleAdmin,
		}

		if err := r.DB.First(new(models.User), "username = ?", support.Username); err == nil {
			// Already have an support
			fmt.Println("Support User Already Configured...")
			return nil
		}
		r.DB.Create(support)
	}
	if err := r.DB.First(new(models.User), "role = ?", models.UserRoleAdmin); err == nil {
		// Already have an admin
		return nil
	}

	fmt.Println("No Admin detected, creating new and printing credentials:")

	password := util.RandomString(16)

	admin := &models.User{
		Username: "admin",
		Password: password,
		Role:     models.UserRoleAdmin,
	}
	r.DB.Create(admin)

	// Print to STDOUT (not c.Log, which would save to file)
	fmt.Printf("\n  ( ͡° ͜ʖ ͡°)  USERNAME: admin  PASSWORD: %s\n\n", password)

	return nil
}
