package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func FindDevices(c *cli.Context) error {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return err
	}
	defer conn.Close()

	address := conn.LocalAddr().String()
	localip := net.ParseIP(address[0:strings.LastIndex(address, ":")]).To4()

	results := make(chan string, 255)
	for n := 1; n <= 254; n++ {
		go func(node int, results chan<- string) {
			remoteip := net.IPv4(localip[0], localip[1], localip[2], byte(node))

			timeout := time.Duration(1 * time.Second)
			client := http.Client{Timeout: timeout}
			res, err := client.Get("http://" + remoteip.String())
			if err == nil && strings.Contains(res.Header.Get("Www-Authenticate"), "rokudev") {
				results <- remoteip.String()
			} else {
				results <- ""
			}
		}(n, results)
	}
	for n := 1; n <= 254; n++ {
		ip := <-results
		if ip != "" {
			fmt.Println(ip)
		}
	}

	return nil
}

func SwitchDevice(c *cli.Context) error {
	rc := NewRC()
	choice, _ := strconv.Atoi(c.Args().First())

	for index, device := range rc.Devices {
		device.Current = false
		if index == choice {
			device.Current = true
		}
	}
	rc.Write()

	ListDevices(c)

	return nil
}

func ListDevices(c *cli.Context) error {
	rc := NewRC()
	if len(rc.Devices) > 0 {
		for index, device := range rc.Devices {
			currentDevice := ""
			if device.Current {
				currentDevice = "current"
			}
			fmt.Printf("%d) %s %s (%s/%s) %s", index, device.Name, device.IP, device.Username, device.Password, currentDevice)
			fmt.Println()
		}
	} else {
		fmt.Println("No devices set up. Use `roku-cli device create NAME IP USERNAME PASSWORD DEFAULT` to create your first device.")
	}

	return nil
}

func CreateDevice(c *cli.Context) error {
	rc := NewRC()
	currentArg, _ := strconv.ParseBool(c.Args().Get(4))

	device := &Device{Name: c.Args().Get(0), IP: c.Args().Get(1), Username: c.Args().Get(2), Password: c.Args().Get(3), Current: currentArg}
	rc.Devices = append(rc.Devices, device)
	rc.Write()

	ListDevices(c)

	return nil
}

func UpdateDevice(c *cli.Context) error {
	rc := NewRC()
	choice, _ := strconv.Atoi(c.Args().First())
	currentArg, _ := strconv.ParseBool(c.Args().Get(4))

	if currentArg {
		for _, device := range rc.Devices {
			device.Current = false
		}
	}

	rc.Devices[choice] = &Device{IP: c.Args().Get(1), Username: c.Args().Get(2), Password: c.Args().Get(3), Current: currentArg}
	rc.Write()

	ListDevices(c)

	return nil
}

func DeleteDevice(c *cli.Context) error {
	rc := NewRC()
	choice, _ := strconv.Atoi(c.Args().First())

	rc.Devices = append(rc.Devices[:choice], rc.Devices[choice+1:]...)
	if rc.CurrentDevice() == nil {
		rc.Devices[0].Current = true
	}
	rc.Write()

	ListDevices(c)

	return nil
}
