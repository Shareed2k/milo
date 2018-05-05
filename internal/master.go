package internal

import (
	"fmt"
	"github.com/soheilhy/cmux"
	"net"
)

type master struct {
	Core
	HttpServer
	MasterServer
}

func NewMaster(c Core) Operator {
	// Init Http server
	return &master{c, NewHttp(c), NewGrpcServer(c)}
}

func (s *master) InitBootstrap() error {
	settings := s.GetSettings().GetOptions()
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", settings.HttpPort))

	if err != nil {
		return err
	}

	// Create a cmux object.
	tcpm := cmux.New(list)

	//grpcL := tcpm.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	grpcL := tcpm.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := tcpm.Match(cmux.HTTP1Fast())

	// Create a grpc server
	go s.getGrpcServer().StartServer(grpcL)

	// Creates a HTTP server
	go s.StartServer(httpL)

	// Start serving!
	return tcpm.Serve()
}

func (s *master) getGrpcServer() GrpcServer {
	return s.MasterServer.(GrpcServer)
}
