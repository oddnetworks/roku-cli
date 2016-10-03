package main

import (
	"fmt"
	// "net"
	// "net/http"
	"os"
	// "strconv"
	// "strings"
	// "time"

	"gopkg.in/urfave/cli.v1"
)

var requiredPaths []string = []string{"manifest", "source"}

func Build(c *cli.Context) error {
	rc, err := NewRC()
	if err != nil {
		return cli.NewExitError("new rc failed: "+err.Error(), 1)
	}

	source := c.String("source")
	if source == "" {
		source = "./"
	}

	for _, path := range requiredPaths {
		if _, err := os.Stat(source + path); os.IsNotExist(err) {
			return cli.NewExitError("Not a valid Roku project. Missing: "+source+path, 1)
		}
	}

	destination := c.String("destination")
	if destination == "" {
		destination = source + "build/"
	}
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		err = os.Mkdir(destination, os.ModePerm)
	}
	if err != nil {
		return cli.NewExitError("Creating destination directory failed: "+err.Error(), 1)
	}

	zip := c.String("zip")
	if zip == "" {
		zip = destination + "channel.zip"
	}

	device := rc.CurrentDevice()
	fmt.Println(source, destination, zip, device)

	return nil
}

func Install(c *cli.Context) error {
	return nil
}
