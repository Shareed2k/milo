package internal

import (
	"fmt"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

type GrpcClient interface {
	Join (in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	ConnectToServer(ip string)
	Close()
}

type client struct {
	Settings
	MasterClient
	connection *grpc.ClientConn
}

func NewGrpcClient(s Settings) GrpcClient {
	return &client{Settings: s}
}

func (s *client) ConnectToServer(ip string) {
	var err error

	s.connection, err = grpc.Dial(fmt.Sprintf("%s:%s", ip, s.GrpcPort), grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	s.MasterClient = NewMasterClient(s.connection)
}

func (s *client) Close() {
	s.connection.Close()
}

func (s *client) Join (in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	return s.MasterClient.Join(context.Background(), in, opts...)
}