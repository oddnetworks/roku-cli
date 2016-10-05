package main

import cli "gopkg.in/urfave/cli.v1"

type flagset struct {
	Choice      int
	Name        string
	Username    string
	Password    string
	IP          string
	Current     bool
	Source      string
	Destination string
	Zip         string
}

var fs flagset

var choiceFlag = cli.IntFlag{
	Name:        "choice, c",
	Usage:       "Choice of Roku device from list",
	Destination: &fs.Choice,
}

var nameFlag = cli.StringFlag{
	Name:        "name, n",
	Usage:       "Name of your Roku device for reference",
	Destination: &fs.Name,
}

var usernameFlag = cli.StringFlag{
	Name:        "username, u",
	Usage:       "Username used to login to with Basic Auth",
	Destination: &fs.Username,
}

var passwordFlag = cli.StringFlag{
	Name:        "password, p",
	Usage:       "Psername used to login to with Basic Auth",
	Destination: &fs.Password,
}

var ipFlag = cli.StringFlag{
	Name:        "ip, i",
	Usage:       "IP address of your Roku device on your local network",
	Destination: &fs.IP,
}

var defaultFlag = cli.BoolFlag{
	Name:        "default, d",
	Usage:       "Set this as the default Roku device to use",
	Destination: &fs.Current,
}

var sourceFlag = cli.StringFlag{
	Name:        "source, src, s",
	Usage:       "Source folder path of your Roku channel, defaults to ./",
	Destination: &fs.Source,
}

var destinationFlag = cli.StringFlag{
	Name:        "destination, dst, d",
	Usage:       "Destination folder path of your Roku channel, defaults to ./build",
	Destination: &fs.Destination,
}

var zipFlag = cli.StringFlag{
	Name:        "zip, z",
	Usage:       "ZIP file name of your Roku channel, defaults to channel.zip",
	Destination: &fs.Zip,
}
