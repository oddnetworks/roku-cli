package commands

import (
	"gopkg.in/urfave/cli.v1"
)

type FlagSet struct {
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

var FS FlagSet

var ChoiceFlag = cli.IntFlag{
	Name:        "choice, c",
	Usage:       "Choice of Roku device from list",
	Destination: &FS.Choice,
}

var NameFlag = cli.StringFlag{
	Name:        "name, n",
	Usage:       "Name of your Roku device for reference",
	Destination: &FS.Name,
}

var UsernameFlag = cli.StringFlag{
	Name:        "username, u",
	Usage:       "Username used to login to with Basic Auth",
	Destination: &FS.Username,
}

var PasswordFlag = cli.StringFlag{
	Name:        "password, p",
	Usage:       "Psername used to login to with Basic Auth",
	Destination: &FS.Password,
}

var IPFlag = cli.StringFlag{
	Name:        "ip, i",
	Usage:       "IP address of your Roku device on your local network",
	Destination: &FS.IP,
}

var DefaultFlag = cli.BoolFlag{
	Name:        "default, d",
	Usage:       "Set this as the default Roku device to use",
	Destination: &FS.Current,
}

var SourceFlag = cli.StringFlag{
	Name:        "source, src, s",
	Usage:       "Source folder path of your Roku channel, defaults to ./",
	Destination: &FS.Source,
}

var DestinationFlag = cli.StringFlag{
	Name:        "destination, dst, d",
	Usage:       "Destination folder path of your Roku channel, defaults to ./build",
	Destination: &FS.Destination,
}

var ZipFlag = cli.StringFlag{
	Name:        "zip, z",
	Usage:       "ZIP file name of your Roku channel, defaults to channel.zip",
	Destination: &FS.Zip,
}
