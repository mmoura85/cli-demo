package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mmoura85/cli-demo/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	//	app details
	app.Name = "test cli app"
	app.Usage = "make an explosive entrance"
	app.Version = "0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Michael",
			Email: "test@example.com",
		},
	}
	app.Copyright = "(c) 2017 "

	app.Commands = []cli.Command{
		cmd.AddCmd,
		//		cmd.RemoveCmd,
		//		cmd.RefreshCmd,
	}

	//	Initialise app
	app.Run(os.Args)
}
