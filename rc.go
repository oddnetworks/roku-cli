package main

import (
	"encoding/json"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
)

type Device struct {
	Name     string `json:"name"`
	IP       string `json:"ip"`
	Username string `json:"username"`
	Password string `json:"password"`
	Current  bool   `json:"current"`
}

type RC struct {
	Devices []*Device `json:"devices"`
	path    string
}

func NewRC() *RC {
	homeDir, _ := homedir.Dir()
	path, _ := homedir.Expand(homeDir)
	path += "/.rokuclirc"

	rc := &RC{
		path: path,
	}

	if _, err := os.Stat(rc.path); os.IsNotExist(err) {
		rc.Write()
	} else {
		rc.Read()
	}

	return rc
}

func (rc *RC) Write() {
	data, _ := json.Marshal(rc)
	ioutil.WriteFile(rc.path, data, os.ModePerm)
}

func (rc *RC) Read() {
	data, _ := ioutil.ReadFile(rc.path)
	json.Unmarshal(data, rc)
}

func (rc *RC) CurrentDevice() *Device {
	for _, device := range rc.Devices {
		if device.Current {
			return device
		}
	}

	return nil
}
