package minion

import (
	"github.com/urfave/cli"
)

func New() cli.Command {
	return cli.Command{
		Name:    "minion",
		Aliases: []string{"m"},
		Usage:   "minion server",
		/*Action: func(c *cli.Context) error {
			fmt.Println("added task: ", c.Args().First())
			return nil
		},*/
		Subcommands: []cli.Command{

		},
	}
}
