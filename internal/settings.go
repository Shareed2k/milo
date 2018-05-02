package internal

import "github.com/urfave/cli"

type Settings interface {
	InitFlags() []cli.Flag
	GetOptions() *settings
}

type settings struct {
	ConfigFilePath string `json:"-"`
	HttpPort       int32  `json:"http_port"`
}

func NewSettings() Settings {
	return &settings{}
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
	}
}

func (s *settings) GetOptions() *settings {
	return s
}
