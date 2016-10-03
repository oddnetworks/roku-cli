package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Roku CLI"
	app.Usage = "Simple CLI for interfacing and developing with a Roku on your local network"

	app.Commands = []cli.Command{
		{
			Name:    "devices",
			Aliases: []string{"d"},
			Usage:   "Manage devices",
			Subcommands: []cli.Command{
				{Name: "find", Aliases: []string{"f"}, Action: FindDevices},
				{Name: "switch", Aliases: []string{"s"}, Action: SwitchDevice,
					Flags: []cli.Flag{
						choiceFlag,
					},
				},
				{Name: "list", Aliases: []string{"l", "ls"}, Action: ListDevices},
				{Name: "create", Aliases: []string{"c"}, Action: CreateDevice,
					Flags: []cli.Flag{
						nameFlag,
						ipFlag,
						usernameFlag,
						passwordFlag,
						defaultFlag,
					},
				},
				{Name: "update", Aliases: []string{"u"}, Action: UpdateDevice,
					Flags: []cli.Flag{
						choiceFlag,
						nameFlag,
						ipFlag,
						usernameFlag,
						passwordFlag,
						defaultFlag,
					},
				},
				{Name: "delete", Aliases: []string{"d", "del", "rm"}, Action: DeleteDevice,
					Flags: []cli.Flag{
						choiceFlag,
					},
				},
			},
		},
		{
			Name:  "install",
			Usage: "Install an app onto the Roku.",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "build",
			Usage: "Build a .zip of the app for submission to the Roku store.",
		},
	}

	app.Run(os.Args)
}
