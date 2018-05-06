package internal

func init() {
	Register("user", func(c Core) (interface{}, error) {
		return NewUserRepository(c), nil
	})
}
