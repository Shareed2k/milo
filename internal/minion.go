package internal

type minion struct {
	Core
}


func NewMinion(c Core) Operator {

	return &minion{c}
}

func (s *minion) InitBootstrap () error {

	return nil
}