package internal

import (
	"fmt"
	"net"
)

type MinionOperator interface {
	InitBootstrap () error
	Close()
}

type minion struct {
	Core
	KeyValueStore
	GrpcClient
	MinionServer
}

func NewMinion(c Core) MinionOperator {
	return &minion{
		Core: c,
		KeyValueStore: NewKeyValueStore(c.GetSettings()),
		GrpcClient: NewGrpcClient(c.GetSettings()),
		MinionServer: NewMinionGrpcServer(c),
	}
}

func (m *minion) InitBootstrap () error {
	settings := m.GetSettings()

	list, err := net.Listen("tcp", fmt.Sprintf(":%s", settings.GrpcPort))

	if err != nil {
		return err
	}

	// Connect to master grpc server
	m.ConnectToServer(settings.MasterAddr)

	return m.getGrpcServer().StartServer(list)
}

func (m *minion) Close() {
	m.GrpcClient.Close()
}

func (m *minion) getGrpcServer() MinionGrpcServer {
	return m.MinionServer.(MinionGrpcServer)
}