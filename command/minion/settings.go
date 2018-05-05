package minion

import "github.com/urfave/cli"

type Settings interface {
	InitFlags() []cli.Flag
	GetOptions() *settings
}

type settings struct {
	Token string `json:"token"`
}

func NewSettings() Settings {
	return &settings{}
}

func (s *settings) InitFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "token, t",
			EnvVar:      "JOIN_TOKEN",
			Usage:       "token to join to master server",
			Destination: &s.Token,
		},
	}
}

func (s *settings) GetOptions() *settings {
	return s
}
