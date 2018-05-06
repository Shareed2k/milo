package internal

import "github.com/urfave/cli"

type Settings interface {
	InitFlags() []cli.Flag
	GetOptions() *settings
}

type settings struct {
	ConfigFilePath  string `json:"-"`
	SupportPassword string `json:"support_password"`
	BindAddr        string `json:"bind_addr"`
	PrivateAddr     string `json:"private_addr"`
	MasterAddr      string `json:"master_addr"`
	HttpPort        int    `json:"http_port"`
	GrpcPort        int    `json:"grpc_port"`
	MasterMode      bool   `json:"master"`
	MinionMode      bool   `json:"minion"`
}

func NewSettings() Settings {
	return &settings{MasterMode: false, MinionMode: true}
}

func (s *settings) InitFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "config, c",
			EnvVar:      "CONFIG",
			Usage:       "config file path",
			Value:       "/etc/milo/milo.json",
			Destination: &s.ConfigFilePath,
		},
		cli.BoolFlag{
			Name:        "master",
			EnvVar:      "MASTER",
			Usage:       "master mode",
			Destination: &s.MasterMode,
		},
		cli.BoolFlag{
			Name:        "minion",
			EnvVar:      "MINION",
			Usage:       "minion mode",
			Destination: &s.MinionMode,
		},
		cli.IntFlag{
			Name:        "http_port",
			EnvVar:      "HTTP_PORT",
			Usage:       "http port",
			Value:       8080,
			Destination: &s.HttpPort,
		},
		cli.IntFlag{
			Name:        "grpc_port",
			EnvVar:      "GRPC_PORT",
			Usage:       "grpc port",
			Value:       8551,
			Destination: &s.GrpcPort,
		},
	}
}

func (s *settings) GetOptions() *settings {
	return s
}
