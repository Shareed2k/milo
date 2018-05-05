package repositories

import (
	"github.com/milo/db/repositories/user"
	"github.com/milo/internal"
)

func init() {
	Register("user", func(c internal.Core) (Repository, error) {
		return user.NewUserRepository(c), nil
	})
}
