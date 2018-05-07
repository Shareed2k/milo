package main

import (
	"fmt"
	"github.com/milo/internal"
	"github.com/urfave/cli"
	"os"
	"github.com/milo/command"
)

func main() {
	app := cli.NewApp()

	app.Name = "Milo"
	app.Usage = "Milo little quite web firewall"
	app.Version = internal.Release

	// init settings struct
	s := internal.NewSettings()

	app.Flags = s.InitFlags()
	app.Commands = command.Map()
	app.Before = func(c *cli.Context) error {
		s.ReadConfig()
		meta := make(map[string]interface{})
		meta["settings"] = s
		app.Metadata = meta
		return nil
	}

	app.Action = func(ctx *cli.Context) {
		c := internal.NewCore(s)
		defer c.OnStop()
	}

	fmt.Printf(
		"%s\ncommit: %s, build time: %s, release: %s\n",
		app.Usage, internal.Commit, internal.BuildTime, internal.Release,
	)

	app.Run(os.Args)
}
