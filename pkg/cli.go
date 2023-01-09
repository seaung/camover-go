package pkg

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	version = "1.0.0"
)

var options = []cli.Flag{}

func run(c *cli.Context) error {
	return nil
}

func Start() {
	author := cli.Author{
		Name: "seaung",
	}

	app := &cli.App{
		Name:    "camover-go",
		Usage:   "CamOver-go is a camera exploitation tool that allows to disclosure network camera admin password.",
		Authors: []*cli.Author{&author},
		Version: version,
		Flags:   options,
		Action:  run,
	}

	if err := app.Run(os.Args); err != nil {
		Errorf(fmt.Sprintf("%v", err))
		os.Exit(1)
	}
}
