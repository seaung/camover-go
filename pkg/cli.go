package pkg

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	version = "1.0.0"
)

func setOutput() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Value:   "console",
		Usage:   "Console print or serialize output.",
	}
}

func setThread() *cli.Int64Flag {
	return &cli.Int64Flag{
		Name:    "thread",
		Aliases: []string{"t"},
		Value:   3,
		Usage:   "Use threads for fast work.",
	}
}

func setInput() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "input",
		Aliases: []string{"i"},
		Value:   "",
		Usage:   "Input file of addresses.",
	}
}

func setAddress() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "address",
		Aliases: []string{"a"},
		Value:   "127.0.0.1",
		Usage:   "Single address.",
	}
}

func setShodan() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "shodan",
		Aliases: []string{"s"},
		Value:   "",
		Usage:   "Shodan API key for exploiting devices over Internet.",
	}
}

var options = []cli.Flag{
	setOutput(),
	setThread(),
	setInput(),
	setAddress(),
	setShodan(),
}

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
