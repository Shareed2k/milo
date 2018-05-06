package internal

func init() {
	Register("user", func(c Core) (Repository, error) {
		return NewUserRepository(c)
	})
	Register("server", func(c Core) (Repository, error) {
		return NewServerRepository(c)
	})
}
