package internal

import (
	"encoding/json"
	"github.com/creasty/defaults"
	"github.com/imdario/mergo"
	"github.com/urfave/cli"
	"os"
)

type Settings struct {
	ConfigFilePath  string `json:"-"`
	SupportPassword string `json:"support_password"`
	BindAddr        string `json:"bind_addr"`
	PrivateAddr     string `json:"private_addr"`
	MasterAddr      string `json:"master_addr"`
	NodeName        string `json:"node_name"`
	HttpPort        string `json:"http_port" default:"8080"`
	GrpcPort        string `json:"grpc_port" default:"8551"`
	MasterMode      bool   `json:"master"`
	MinionMode      bool   `json:"minion"`
}

func NewSettings() Settings {
	return Settings{MasterMode: false, MinionMode: false}
}

func (s *Settings) InitFlags() []cli.Flag {
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
		cli.StringFlag{
			Name:        "http_port",
			EnvVar:      "HTTP_PORT",
			Usage:       "http port",
			Destination: &s.HttpPort,
		},
		cli.StringFlag{
			Name:        "grpc_port",
			EnvVar:      "GRPC_PORT",
			Usage:       "grpc port",
			Destination: &s.GrpcPort,
		},
	}
}

func (s *Settings) ReadConfig() error {
	// Set Settings from config file
	if s.ConfigFilePath != "" {
		var configFileSettings Settings
		configFile, err := os.Open(s.ConfigFilePath)
		defer configFile.Close()

		if err != nil {
			return err
		}
		if err := json.NewDecoder(configFile).Decode(&configFileSettings); err != nil {
			return err
		}

		// Merge in command line settings (which overwrite respective config file settings)
		if err := mergo.Merge(s, configFileSettings); err != nil {
			return err
		}

		// Set Default Settings with struct tags
		if err := defaults.Set(s); err != nil {
			return err
		}
	}

	return nil
}
