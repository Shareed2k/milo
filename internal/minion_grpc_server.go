package internal

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"github.com/milo/ipaddr"
	"fmt"
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
	ports := []*StatsResponse_Process{}
	for _, proto := range []string{"tcp", "tcp6", "udp", "udp6"} {
		for _, p := range netstat(proto) {
			// Check STATE to show only Listening connections
			if p.State == "LISTEN" {
				ports = append(ports, &StatsResponse_Process{
					Ip:          p.Ip,
					Port:        p.Port,
					State:       p.State,
					ProcessName: p.Name,
					User:        p.User,
					Pid:         p.Pid,
					Proto:       proto,
				})
			}
		}
	}

	ports2 := []*StatsResponse_Process{}
	for _, proto := range []string{"tcp", "udp6"} {
		for _, p := range netstat(proto) {
			// Check STATE to show only Listening connections
			if p.State == "LISTEN" {
				ports2 = append(ports2, &StatsResponse_Process{
					Ip:          p.Ip,
					Port:        p.Port,
					State:       p.State,
					ProcessName: p.Name,
					User:        p.User,
					Pid:         p.Pid,
					Proto:       proto,
				})
			}
		}
	}

	ips, err := ipaddr.GetPrivateIPv4()

	if diff := Equal(ports, ports2); diff != nil {
		fmt.Println(diff)
	}

	fmt.Println(ips)

	return &StatsResponse{
		Processes: ports,
	}, err
}
