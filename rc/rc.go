package rc

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
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

func LoadRC() (*RC, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	path, err := homedir.Expand(homeDir)
	if err != nil {
		return nil, err
	}
	path += "/.rokuclirc"

	rc := &RC{
		path: path,
	}

	if _, err = os.Stat(rc.path); os.IsNotExist(err) {
		err = rc.Write()
	} else {
		err = rc.Read()
	}

	if err != nil {
		return nil, err
	}

	return rc, nil
}

func (rc *RC) Write() error {
	data, err := json.Marshal(rc)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(rc.path, data, os.ModePerm)
	return err
}

func (rc *RC) Read() error {
	data, err := ioutil.ReadFile(rc.path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, rc)
	return err
}

func (rc *RC) CurrentDevice() *Device {
	for _, device := range rc.Devices {
		if device.Current {
			return device
		}
	}

	return nil
}
