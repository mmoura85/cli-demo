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
			Name: "name, n",
			//			Value: "",
			Usage: "persons name",
		},
		{
			Name:  "phone, ph",
			Usage: "phone number",
		},
		{
			Name:  "mob, m",
			Usage: "mob number",
		},
	},
}

func addAction(c *cli.Context) {

}
