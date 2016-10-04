package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

var requiredPaths []string = []string{"manifest", "source"}
var allowedPaths []string = []string{"manifest", "source", "images", "components"}

func Build(c *cli.Context) error {
	source := fs.Source
	if source == "" {
		source = "./"
	}

	for _, required := range requiredPaths {
		if _, err := os.Stat(source + required); os.IsNotExist(err) {
			return cli.NewExitError("Not a valid Roku project. Missing: "+source+required, 1)
		}
	}

	destination := fs.Destination
	if destination == "" {
		destination = source + "build/"
	}
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		err = os.Mkdir(destination, os.ModePerm)
	}

	zipName := fs.Zip
	if zipName == "" {
		zipName = destination + "channel.zip"
	} else {
		zipName = destination + zipName
	}

	zipFile, err := os.Create(zipName)
	if err != nil {
		return cli.NewExitError("Zip file could not be created: "+err.Error(), 1)
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	baseDir := filepath.Base(source)
	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, allowed := range allowedPaths {
			if strings.Contains(path, allowed) {
				header, err := zip.FileInfoHeader(info)
				if err != nil {
					return err
				}

				if baseDir != "" {
					header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
				}

				if info.IsDir() {
					header.Name += "/"
				} else {
					header.Method = zip.Deflate
				}

				writer, err := archive.CreateHeader(header)
				if err != nil {
					return err
				}

				if info.IsDir() {
					return nil
				}

				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				_, err = io.Copy(writer, file)
				return err
			}
		}

		return err
	})

	fmt.Println("Zip created: ", zipName)

	return nil
}

func Install(c *cli.Context) error {
	return nil
}
