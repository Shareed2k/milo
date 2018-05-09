package internal

import (
	"github.com/milo/db/models"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"github.com/pkg/errors"
)

type GrpcServer interface {
	StartServer(l net.Listener) error
}

type server struct {
	Core
}

func NewGrpcServer(c Core) MasterServer {
	return &server{
		Core:   c,
	}
}

func (s *server) StartServer(l net.Listener) error {
	server := grpc.NewServer()
	RegisterMasterServer(server, s)

	return server.Serve(l)
}

func (s *server) Join(ctx context.Context, in *JoinRequest) (*JoinResponse, error) {
	model := &models.Server{
		PrivateIp:   in.GetMinion().GetPrivateAddr(),
		PublicIp:    in.GetMinion().PublicAddr,
		Description: "test test description",
	}

	repo := s.GetMaster().GetServerRepository()

	if _, err := repo.Create(model); err != nil {
		return nil, errors.New("Minion is already registered")
	}

	return &JoinResponse{
		Uuid:    model.Uuid,
		Message: model.Description,
	}, nil
}

func (s *server) GetRuleList(ctx context.Context, in *RuleListRequest) (*RuleListResponse, error) {
	return nil, nil
}