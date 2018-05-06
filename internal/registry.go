package internal

import (
	"fmt"
	"errors"
)

type Repository interface {

}

// Factory is a function that returns a new instance of a Repository.
type Factory func(s Core) (Repository, error)

// registry has an entry for each available repository,
// This should be populated at package init() time via Register().
var registry map[string]Factory

// Register adds a new repository to the registry.
func Register(name string, fn Factory) {
	if registry == nil {
		registry = make(map[string]Factory)
	}

	if registry[name] != nil {
		panic(fmt.Errorf("Repository %q is already registered", name))
	}
	registry[name] = fn
}

func Map(c Core) map[string]Repository {
	m := make(map[string]Repository)
	for name, fn := range registry {
		thisFn := fn
		m[name] = func() (Repository, error) {
			return thisFn(c)
		}
	}
	return m
}

func CreateRepository(name string, c Core) (Repository, error) {
	engineFactory, ok := registry[name]
	if !ok {
		// Factory has not been registered.
		// Make a list of all available datastore factories for logging.
		return nil, errors.New(fmt.Sprintf("Invalid Repository name: %s", name))
	}

	// Run the factory with the configuration.
	return engineFactory(c)
}