package internal

import (
	"fmt"
	"github.com/milo/db/models"
	"github.com/soheilhy/cmux"
	"net"
)

type MasterOperator interface {
	GetDatabase() *Database
}

type master struct {
	Core
	HttpServer
	MasterServer
	*Database
}

func NewMaster(c Core) Operator {
	// Init Http, Grpc server
	return &master{
		c,
		NewHttp(c),
		NewGrpcServer(c),
		NewDatabase(c.GetSettings()),
	}
}

func (m *master) InitBootstrap() error {
	settings := m.GetSettings().GetOptions()
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", settings.Port))

	if err != nil {
		return err
	}

	// Run migration
	m.AutoMigrate(&models.User{})

	// Create a cmux object.
	tcpm := cmux.New(list)

	//grpcL := tcpm.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	grpcL := tcpm.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := tcpm.Match(cmux.HTTP1Fast())

	// Create a grpc server
	go m.getGrpcServer().StartServer(grpcL)

	// Creates a HTTP server
	go m.StartServer(httpL)

	// Start serving!
	return tcpm.Serve()
}

func (m *master) GetDatabase() *Database {
	return m.Database
}

func (m *master) getGrpcServer() GrpcServer {
	return m.MasterServer.(GrpcServer)
}
