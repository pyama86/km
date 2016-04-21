package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/pyama86/km/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "attach",
		Usage:  "",
		Action: command.CmdAttach,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "detach",
		Usage:  "",
		Action: command.CmdDetach,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
