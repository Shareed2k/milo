package minion

import (
	"errors"
	"fmt"
	"github.com/milo/internal"
	"github.com/milo/ipaddr"
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
					settings := config.(internal.Settings)

					client := internal.NewGrpcClient(settings)
					client.ConnectToServer(c.Args().First())

					store := internal.NewKeyValueStore(settings)

					privateAddrs := []*internal.JoinRequest_Addr{}
					publicAddrs := []*internal.JoinRequest_Addr{}
					publicAddrs = append(publicAddrs, &internal.JoinRequest_Addr{
						Ip: "123.89.89.7",
					})

					ips, err := ipaddr.GetPrivateIPv4()

					if err != nil {
						println(err.Error())
						return err
					}

					for _, ip := range ips {
						privateAddrs = append(privateAddrs, &internal.JoinRequest_Addr{
							Ip: string(ip.IP),
						})
					}

					request := &internal.JoinRequest{
						Token: s.GetOptions().Token,
						Minion: &internal.JoinRequest_Minion{
							PrivateAddrs: privateAddrs,
							PublicAddrs:  publicAddrs,
						},
					}

					fmt.Printf("added task: %s, token: %s\n", c.Args().First(), s.GetOptions().Token)

					response, err := client.Join(request)

					if err != nil {
						println(err.Error())
						return err
					}

					store.Set("uuid", []byte(response.Uuid))
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
