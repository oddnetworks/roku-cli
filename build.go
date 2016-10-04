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

	// Verify source folder contains required Roku files and folders
	for _, required := range requiredPaths {
		verifyPath := filepath.Join(source, required)
		if _, err := os.Stat(verifyPath); os.IsNotExist(err) {
			return cli.NewExitError("Not a valid Roku project. Missing: "+verifyPath, 1)
		}
	}

	fmt.Println("Building from path:", source)

	destination := fs.Destination
	if destination == "" {
		destination = filepath.Join(source, "build")
	}

	// Make the destination folder if it doesn't exist
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		err = os.Mkdir(destination, os.ModePerm)
	}

	zipName := fs.Zip
	if zipName == "" {
		zipName = filepath.Join(destination, "channel.zip")
	} else {
		zipName = filepath.Join(destination, zipName)
	}

	// Make a new file handler and zip archive
	zipFile, err := os.Create(zipName)
	if err != nil {
		return cli.NewExitError("Zip file could not be created: "+err.Error(), 1)
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	// Walk the source path and add each path to the archive
	baseDir := filepath.Base(source)
	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, allowed := range allowedPaths {
			if strings.Contains(path, allowed) {
				header, err := zip.FileInfoHeader(info)
				if err != nil {
					return err
				}

				header.Name = strings.TrimPrefix(path, fs.Source+"/")

				header.Method = zip.Store
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
	if err != nil {
		return cli.NewExitError("Error zipping: "+err.Error(), 1)
	}

	fmt.Println("Build complete:", zipName)

	return nil
}

func Install(c *cli.Context) error {
	return nil
}
