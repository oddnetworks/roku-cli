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
				{Name: "switch", Aliases: []string{"s"}, Action: SwitchDevice, After: ListDevices,
					Flags: []cli.Flag{
						choiceFlag,
					},
				},
				{Name: "list", Aliases: []string{"l", "ls"}, Action: ListDevices},
				{Name: "create", Aliases: []string{"c"}, Action: CreateDevice, After: ListDevices,
					Flags: []cli.Flag{
						nameFlag,
						ipFlag,
						usernameFlag,
						passwordFlag,
						defaultFlag,
					},
				},
				{Name: "update", Aliases: []string{"u"}, Action: UpdateDevice, After: ListDevices,
					Flags: []cli.Flag{
						choiceFlag,
						nameFlag,
						ipFlag,
						usernameFlag,
						passwordFlag,
						defaultFlag,
					},
				},
				{Name: "delete", Aliases: []string{"d", "del", "rm"}, Action: DeleteDevice, After: ListDevices,
					Flags: []cli.Flag{
						choiceFlag,
					},
				},
			},
		},
		{
			Name:    "install",
			Aliases: []string{"i"},
			Flags:   []cli.Flag{sourceFlag, destinationFlag, zipFlag},
			Usage:   "Install an app onto the Roku.",
			Before:  EnsurePaths,
			Action:  Install,
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Flags:   []cli.Flag{sourceFlag, destinationFlag, zipFlag},
			Usage:   "Build a .zip of the app for submission to the Roku store.",
			Before:  EnsurePaths,
			Action:  Build,
		},
	}

	app.Run(os.Args)
}
