package internal

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

type MinionGrpcServer interface {
	StartServer(l net.Listener) error
}

type minionServer struct {
	Core
}

func NewMinionGrpcServer(c Core) MinionServer {
	return &minionServer{
		Core: c,
	}
}

func (s *minionServer) StartServer(l net.Listener) error {
	server := grpc.NewServer()
	RegisterMinionServer(server, s)

	return server.Serve(l)
}

func (s *minionServer) PassRule(ctx context.Context, in *RuleRequest) (*RuleResponse, error) {
	return nil, nil
}

func (s *minionServer) GetStats(ctx context.Context, in *StatsRequest) (*StatsResponse, error) {
	return nil, nil
}
