package internal

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"github.com/ugorji/go/codec"
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
	nums := map[int64]int64{}
	for _, proto := range []string{"tcp", "tcp6", "udp", "udp6"} {
		for _, p := range netstat(proto) {
			// Check STATE to show only Listening connections
			if p.State == "LISTEN" {
				nums[p.Port] = p.Port
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

	// TODO: need to create ticker and create task to get post and decode to byte and save in badger
	// on request from master return response from budger

	kv := s.GetMinion().GetKeyValueStore()
	val, err := kv.Get("ports_list")

	var mh codec.MsgpackHandle

	// we have last value of ports
	last_ports := map[int64]int64{}
	if err == nil {
		dec := codec.NewDecoderBytes(val, &mh)
		err = dec.Decode(&last_ports)

		if diff := Equal(val, last_ports); diff != nil {
			fmt.Println(diff)
		}
	} else { // fresh need to save
		raw := []byte{}
		enc := codec.NewEncoderBytes(&raw, &mh)
		err = enc.Encode(nums)

		kv.Set("ports_list", raw)
	}

	ips, err := ipaddr.GetPrivateIPv4()

	fmt.Println(ips)

	return &StatsResponse{
		Processes: ports,
	}, err
}
