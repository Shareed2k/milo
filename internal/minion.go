package internal

type MinionOperator interface {
	InitBootstrap () error
	Close()
}

type minion struct {
	Core
	KeyValueStore
	GrpcClient
}

func NewMinion(c Core) MinionOperator {
	return &minion{
		Core: c,
		KeyValueStore: NewKeyValueStore(c.GetSettings()),
		GrpcClient: NewGrpcClient(c.GetSettings()),
	}
}

func (m *minion) InitBootstrap () error {
	settings := m.GetSettings().GetOptions()

	// Connect to master grpc server
	m.ConnectToServer(settings.MasterAddr)

	return nil
}

func (m *minion) Close() {
	m.GrpcClient.Close()
}