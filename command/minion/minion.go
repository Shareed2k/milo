package minion

import (
	"github.com/urfave/cli"
	"fmt"
	"github.com/milo/internal"
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
					fmt.Println(st.GetOptions())
					fmt.Printf("added task: %s, token: %s", c.Args().First(), s.GetOptions().Token)
					return nil
				},
			},
		},
	}
}
