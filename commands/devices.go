package commands

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"gopkg.in/urfave/cli.v1"

	"github.com/oddnetworks/roku-cli/rc"
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
			client := &http.Client{Timeout: timeout}
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
	config, err := rc.LoadRC()
	if err != nil {
		return cli.NewExitError("new rc failed: "+err.Error(), 1)
	}

	for index, device := range config.Devices {
		device.Current = false
		if index == FS.Choice {
			device.Current = true
		}
	}
	config.Write()

	return nil
}

func ListDevices(c *cli.Context) error {
	config, err := rc.LoadRC()
	if err != nil {
		return cli.NewExitError("new rc failed: "+err.Error(), 1)
	}
	if len(config.Devices) > 0 {
		for index, device := range config.Devices {
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
	config, err := rc.LoadRC()
	if err != nil {
		return cli.NewExitError("new rc failed: "+err.Error(), 1)
	}

	if FS.Name == "" {
		return cli.NewExitError("Missing --name flag.", 1)
	}
	if FS.IP == "" {
		return cli.NewExitError("Missing --ip flag.", 1)
	}
	if FS.Username == "" {
		return cli.NewExitError("Missing --username flag.", 1)
	}
	if FS.Password == "" {
		return cli.NewExitError("Missing --password flag.", 1)
	}

	device := &rc.Device{Name: FS.Name, IP: FS.IP, Username: FS.Username, Password: FS.Password, Current: FS.Current}
	config.Devices = append(config.Devices, device)
	config.Write()

	return nil
}

func UpdateDevice(c *cli.Context) error {
	config, err := rc.LoadRC()
	if err != nil {
		return cli.NewExitError("new rc failed: "+err.Error(), 1)
	}

	if FS.Current {
		for _, device := range config.Devices {
			device.Current = false
		}
	}

	if FS.Choice+1 > len(config.Devices) {
		return cli.NewExitError("invalid device number to update", 1)
	}

	device := config.Devices[FS.Choice]
	if FS.Name != "" {
		device.Name = FS.Name
	}
	if FS.IP != "" {
		device.IP = FS.IP
	}
	if FS.Username != "" {
		device.Username = FS.Username
	}
	if FS.Password != "" {
		device.Password = FS.Password
	}
	device.Current = FS.Current

	config.Write()

	return nil
}

func DeleteDevice(c *cli.Context) error {
	config, err := rc.LoadRC()
	if err != nil {
		return cli.NewExitError("new rc failed: "+err.Error(), 1)
	}

	if FS.Choice+1 > len(config.Devices) {
		return cli.NewExitError("invalid device number to delete", 1)
	}

	config.Devices = append(config.Devices[:FS.Choice], config.Devices[FS.Choice+1:]...)
	if config.CurrentDevice() == nil && len(config.Devices) > 0 {
		config.Devices[0].Current = true
	}
	config.Write()

	return nil
}
