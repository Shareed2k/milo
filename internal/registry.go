package internal

import (
	"fmt"
)

// Factory is a function that returns a new instance of a Repository.
type Factory func(s Core) (interface{}, error)

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

func Map(c Core) map[string]interface{} {
	m := make(map[string]interface{})
	for name, fn := range registry {
		thisFn := fn
		m[name] = func() (interface{}, error) {
			return thisFn(c)
		}
	}
	return m
}