package command

import (
	"github.com/milo/command/minion"
	"github.com/urfave/cli"
)

func init() {
	Register("minion", func() (cli.Command, error) {
		return minion.New(), nil
	})
}
