package internal

import (
	"fmt"
	"github.com/milo/db/models"
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
	return &server{
		Core:   c,
	}
}

func (s *server) StartServer(l net.Listener) {
	server := grpc.NewServer()
	RegisterMasterServer(server, s)

	server.Serve(l)
}

func (s *server) Join(ctx context.Context, in *JoinRequest) (*JoinResponse, error) {
	fmt.Println(in)

	model := &models.Server{
		PrivateIp:   in.GetMinion().GetPrivateAddr(),
		PublicIp:    in.GetMinion().PublicAddr,
		Description: "test test description",
	}

	repo := s.GetMaster().GetServerRepository()
	repo.Create(model)

	return &JoinResponse{
		Uuid:    model.Uuid,
		Message: model.Description,
	}, nil
}

func (s *server) StreamRule(in *Rule, str Master_StreamRuleServer) error {
	return nil
}

func (s *server) GetListRule(ctx context.Context, in *FetchRequest) (*ListRules, error) {
	return nil, nil
}
