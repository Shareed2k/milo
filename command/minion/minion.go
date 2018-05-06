package minion

import (
	"fmt"
	"github.com/milo/internal"
	"github.com/urfave/cli"
)

func New(st internal.Settings) cli.Command {
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
					fmt.Println(st)

					client := internal.NewGrpcClient(st)
					client.ConnectToServer(c.Args().First())

					store := internal.NewKeyValueStore(st)

					request := &internal.JoinRequest{
						Token: s.GetOptions().Token,
						Minion: &internal.JoinRequest_Minion{
							PrivateAddr: st.GetOptions().BindAddr,
							PublicAddr: "122.45.65.12",
						},
					}

					fmt.Printf("added task: %s, token: %s\n", c.Args().First(), s.GetOptions().Token)

					response, err := client.Join(request)

					if err != nil {
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
