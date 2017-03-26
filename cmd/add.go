package cmd

import (
	"github.com/urfave/cli"
)

var AddCmd = cli.Command{
	Name: "add",
	//	Aliases: []string{"a"},
	Usage:  "add a new value",
	Action: addAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "test, t",
			Value: "",
			Usage: "add a test value",
		},
	},
}

func addAction(c *cli.Context) {

}
