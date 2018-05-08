package internal

import (
	"fmt"
	"github.com/milo/db/models"
	"github.com/soheilhy/cmux"
	"net"
)

type MasterOperator interface {
	GetDatabase() *Database
	GetServerRepository() ServerRepository
	InitBootstrap () error
}

type master struct {
	Core
	HttpServer
	MasterServer
	*Database
	userRepo UserRepository
	serverRepo ServerRepository
}

func NewMaster(c Core) MasterOperator {
	// Init Http, Grpc server
	return &master{
		Core: c,
		HttpServer: NewHttp(c),
		MasterServer: NewGrpcServer(c),
		Database: NewDatabase(c.GetSettings()),
	}
}

func (m *master) InitBootstrap() error {
	settings := m.GetSettings()
	list, err := net.Listen("tcp", fmt.Sprintf(":%s", settings.GrpcPort))
	httpList, err := net.Listen("tcp", fmt.Sprintf(":%s", settings.HttpPort))

	if err != nil {
		return err
	}

	// Run migration
	m.AutoMigrate(&models.User{})
	m.AutoMigrate(&models.Server{})
	m.AutoMigrate(&models.DataCenter{})

	// Create admin user
	userRepo, _ := CreateRepository("user", m.Core)
	serverRepo, _ := CreateRepository("server", m.Core)
	m.userRepo = userRepo.(UserRepository)
	m.serverRepo = serverRepo.(ServerRepository)

	if err := m.userRepo.DetectOrCreateAdmin(); err != nil {
		return err
	}

	// Create a cmux object.
	tcpm := cmux.New(list)

	//grpcL := tcpm.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	grpcL := tcpm.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	//httpL := tcpm.Match(cmux.HTTP1Fast())

	// Create a grpc server
	go m.getGrpcServer().StartServer(grpcL)

	// Creates a HTTP server
	go m.StartServer(httpList)

	// Start serving!
	return tcpm.Serve()
}

func (m *master) GetDatabase() *Database {
	return m.Database
}

func (m *master) getGrpcServer() GrpcServer {
	return m.MasterServer.(GrpcServer)
}

func (m *master) GetServerRepository() ServerRepository {
	return m.serverRepo
}

func (m *master) Close() {

}