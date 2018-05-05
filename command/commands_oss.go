package command

import (
	"github.com/milo/command/minion"
	"github.com/urfave/cli"
	"github.com/milo/internal"
)

func init() {
	Register("minion", func(s internal.Settings) (cli.Command, error) {
		return minion.New(s), nil
	})
}
