package minion

import (
	"errors"
	"fmt"
	"github.com/milo/internal"
	"github.com/urfave/cli"
)

func New() cli.Command {
	s := NewSettings()

	return cli.Command{
		Name:    "minion",
		Aliases: []string{"m"},
		Usage:   "minion server",
		Subcommands: []cli.Command{
			{
				Name:  "join",
				Usage: "join to master",
				Flags: s.InitFlags(),
				Action: func(c *cli.Context) error {
					var config interface{}
					var ok bool
					if config, ok = c.App.Metadata["settings"]; !ok {
						return errors.New("settings is missing")
					}
					settings := config.(internal.Settings).GetOptions()

					client := internal.NewGrpcClient(settings)
					client.ConnectToServer(c.Args().First())

					store := internal.NewKeyValueStore(settings)

					request := &internal.JoinRequest{
						Token: s.GetOptions().Token,
						Minion: &internal.JoinRequest_Minion{
							PrivateAddr: settings.GetOptions().BindAddr,
							PublicAddr:  "122.45.65.12",
						},
					}

					fmt.Printf("added task: %s, token: %s\n", c.Args().First(), s.GetOptions().Token)

					response, err := client.Join(request)

					if err != nil {
						println(err.Error())
						return err
					}

					store.Set("uuid", response.Uuid)
					val, err := store.Get("uuid")

					if err != nil {
						return err
					}

					fmt.Printf("badget db: %s\n", val)

					fmt.Println(response)

					return nil
				},
			},
		},
	}
}
