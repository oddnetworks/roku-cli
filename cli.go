package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"

	"github.com/oddnetworks/roku-cli/commands"
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
				{Name: "find", Aliases: []string{"f"}, Action: commands.FindDevices},
				{Name: "switch", Aliases: []string{"s"}, Action: commands.SwitchDevice, After: commands.ListDevices,
					Flags: []cli.Flag{commands.ChoiceFlag},
				},
				{Name: "list", Aliases: []string{"l", "ls"}, Action: commands.ListDevices},
				{Name: "create", Aliases: []string{"c"}, Action: commands.CreateDevice, After: commands.ListDevices,
					Flags: []cli.Flag{
						commands.NameFlag,
						commands.IPFlag,
						commands.UsernameFlag,
						commands.PasswordFlag,
						commands.DefaultFlag,
					},
				},
				{Name: "update", Aliases: []string{"u"}, Action: commands.UpdateDevice, After: commands.ListDevices,
					Flags: []cli.Flag{
						commands.ChoiceFlag,
						commands.NameFlag,
						commands.IPFlag,
						commands.UsernameFlag,
						commands.PasswordFlag,
						commands.DefaultFlag,
					},
				},
				{Name: "delete", Aliases: []string{"d", "del", "rm"}, Action: commands.DeleteDevice, After: commands.ListDevices,
					Flags: []cli.Flag{commands.ChoiceFlag},
				},
			},
		},
		{
			Name:    "install",
			Aliases: []string{"i"},
			Flags:   []cli.Flag{commands.SourceFlag, commands.DestinationFlag, commands.ZipFlag},
			Usage:   "Install an app onto the Roku.",
			Before:  commands.EnsurePaths,
			Action:  commands.Install,
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Flags:   []cli.Flag{commands.SourceFlag, commands.DestinationFlag, commands.ZipFlag},
			Usage:   "Build a .zip of the app for submission to the Roku store.",
			Before:  commands.EnsurePaths,
			Action:  commands.Build,
		},
	}

	app.Run(os.Args)
}
