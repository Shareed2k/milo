package internal

func init() {
	Register("user", func(c Core) (Repository, error) {
		return NewUserRepository(c)
	})
}
