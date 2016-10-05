package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

type authorization struct {
	Username, Password, Realm, NONCE, QOP, Opaque, Algorithm string
}

func Install(c *cli.Context) error {
	Build(c)

	// Open the rc file and get the current device
	rc, _ := NewRC()
	device := rc.CurrentDevice()

	// Open the zip file
	zip, err := os.Open(fs.Zip)
	if err != nil {
		return cli.NewExitError("Error reading zip file: "+err.Error(), 1)
	}
	defer zip.Close()

	// Build a form and add the zip binary file
	form := &bytes.Buffer{}
	writer := multipart.NewWriter(form)
	part, err := writer.CreateFormFile("archive", filepath.Base(fs.Zip))
	if err != nil {
		return cli.NewExitError("Error attaching zip file: "+err.Error(), 1)
	}
	_, err = io.Copy(part, zip)

	writer.WriteField("mysubmit", "Install")
	writer.Close()

	// Simple auth struct
	auth := &authorization{"rokudev", device.Password, "rokudev", "", "auth", "", ""}

	// Begin building HTTP Digest Auth
	login := strings.Join([]string{auth.Username, auth.Realm, auth.Password}, ":")
	h := md5.New()
	io.WriteString(h, login)
	loginHash := hex.EncodeToString(h.Sum(nil))

	action := strings.Join([]string{"POST", "/plugin_install"}, ":")
	h = md5.New()
	io.WriteString(h, action)
	actionHash := hex.EncodeToString(h.Sum(nil))

	nc_str := fmt.Sprintf("%08x", 3)
	hnc := "MTM3MDgw"

	responseDigest := fmt.Sprintf("%s:%s:%s:%s:%s:%s", loginHash, auth.NONCE, nc_str, hnc, auth.QOP, actionHash)
	h = md5.New()
	io.WriteString(h, responseDigest)
	responseDigest = hex.EncodeToString(h.Sum(nil))

	digest := "username=\"%s\", realm=\"%s\", nonce=\"%s\", uri=\"%s\", response=\"%s\""
	digest = fmt.Sprintf(digest, auth.Username, auth.Realm, auth.NONCE, "/plugin_install", responseDigest)
	if auth.Opaque != "" {
		digest += fmt.Sprintf(", opaque=\"%s\"", auth.Opaque)
	}
	if auth.QOP != "" {
		digest += fmt.Sprintf(", qop=\"%s\", nc=%s, cnonce=\"%s\"", auth.QOP, nc_str, hnc)
	}
	if auth.Algorithm != "" {
		digest += fmt.Sprintf(", algorithm=\"%s\"", auth.Algorithm)
	}

	// Post the form with the digest auth to the Roku device
	req, err := http.NewRequest("POST", "http://"+device.IP+"/plugin_install", form)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Digest "+digest)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return cli.NewExitError("Error installing build: "+err.Error(), 1)
	} else {
		if res.StatusCode == 401 {
			return cli.NewExitError("Error installing build: Username/Password incorrect for "+device.IP, 1)
		}

		// Parse the HTML for the Roku message
		resBody, _ := ioutil.ReadAll(res.Body)
		body := string(resBody)
		messageIndex := strings.Index(body, "Roku.Message")
		scriptIndex := strings.LastIndex(body, "Render")
		message := body[messageIndex+15 : scriptIndex-10]
		triggers := strings.Split(message, ".")
		content := strings.Split(triggers[1], "', '")
		fmt.Println("Install complete:", device.Name, device.IP)
		fmt.Println("Roku Response:", "\""+content[1]+"\"")
	}

	return nil
}
