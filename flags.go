package main

import cli "gopkg.in/urfave/cli.v1"

type flagset struct {
	Choice   int
	Name     string
	Username string
	Password string
	IP       string
	Current  bool
}

var fs flagset

var choiceFlag = cli.IntFlag{
	Name:        "choice",
	Usage:       "choice of roku device from list",
	Destination: &fs.Choice,
}

var nameFlag = cli.StringFlag{
	Name:        "name",
	Usage:       "name of roku device",
	Destination: &fs.Name,
}

var usernameFlag = cli.StringFlag{
	Name:        "username",
	Usage:       "username of roku device",
	Destination: &fs.Username,
}

var passwordFlag = cli.StringFlag{
	Name:        "password",
	Usage:       "password for roku device",
	Destination: &fs.Password,
}

var ipFlag = cli.StringFlag{
	Name:        "ip",
	Usage:       "ip address of roku device",
	Destination: &fs.IP,
}

var defaultFlag = cli.BoolFlag{
	Name:        "default",
	Usage:       "is default roku device to use",
	Destination: &fs.Current,
}
