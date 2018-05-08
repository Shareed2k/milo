package internal

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type MinionGrpcClient interface {
	GetStats(in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error)
	ConnectToServer(ip, port string)
	Close()
}

type minionGrpcClient struct {
	settings Settings
	MinionClient
	connection *grpc.ClientConn
}

func NewMinionGrpcClient(s Settings) MinionGrpcClient {
	return &minionGrpcClient{settings: s}
}

func (s *minionGrpcClient) ConnectToServer(ip, port string) {
	var err error

	s.connection, err = grpc.Dial(fmt.Sprintf("%s:%s", ip, port), grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	s.MinionClient = NewMinionClient(s.connection)
}

func (s *minionGrpcClient) Close() {
	s.connection.Close()
}

func (s *minionGrpcClient) GetStats(in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error) {
	return s.MinionClient.GetStats(context.Background(), in, opts...)
}

//PassRule(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*RuleResponse, error)
