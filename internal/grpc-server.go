package internal

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer interface {
	StartServer(l net.Listener)
}

type server struct {
	Core
}

func NewGrpcServer(c Core) MasterServer {
	return &server{c}
}

func (s *server) StartServer(l net.Listener) {
	server := grpc.NewServer()
	RegisterMasterServer(server, s)

	server.Serve(l)
}

func (s *server) Join(ctx context.Context, in *JoinRequest) (*JoinResponse, error) {
	return nil, nil
}

func (s *server) StreamRule(in *Rule, str Master_StreamRuleServer) error {
	return nil
}

func (s *server) GetListRule(ctx context.Context, in *FetchRequest) (*ListRules, error) {
	return nil, nil
}