# Roku CLI

A cross platform Roku CLI written in Go.

## Features

- Manages multiple Roku devices during development
- Debug Roku output from Telnet ports (TODO)
- Builds and installs apps onto Roku hardware (TODO)
- Builds apps for submission to the Roku Channel Store (TODO)
- Scans local network to find all available Roku devices

## Installing

1. Install Go
2. `go install github.com/oddnetworks/roku-cli`

## Usage

`$ roku-cli COMMAND SUBCOMMAND ARGS`

### Command: `devices|d`

All of the `devices` commands work against your `$HOME/.rokuclirc` file to store IP address, usernames, and passwords to make working with multiple devices easier.

You can set a **current** devices and switch between them so that commands like `install` will always work with the currently active device.

#### Subcommand: `find|f`

Scans the local subnet for all available Roku devices.

```
$ roku-cli devices find
192.168.0.12
192.168.0.100
```

#### Subcommand: `list|l|ls`

List all Roku devices set in your `$HOME/.rokuclirc` file.

```
$ roku-cli devices list
0) stick 192.168.0.13 (rokudev/testing12345)
1) roku-3 192.168.0.100 (admin/demo) current
```

#### Subcommand: `switch|s`

Switch Roku devices to set a current one for interacting with.

```
$ roku-cli devices switch -choice=0
0) stick 192.168.0.13 (rokudev/testing12345) current
1) roku-3 192.168.0.100 (admin/demo)
```

#### Subcommand: `create|c`

Create a new Roku device in your `$HOME/.rokuclirc` file. All flags are required to create a device.

```
$ roku-cli devices create --name="roku-2" --ip="10.0.0.12" --username=rokudev --password=roku4life --default=true
0) stick 192.168.0.13 (rokudev/testing12345)
1) roku-3 192.168.0.100 (admin/demo)
2) roku-2 10.0.0.12 (rokudev/roku4life) current
```

#### Subcommand: `update|u`

Update a Roku device in your `$HOME/.rokuclirc` file. Only flag values that are passed in will be updated. Others will be left alone.

```
$ roku-cli devices update --choice=1 --name="roku-3" --ip="10.0.0.100"
0) stick 192.168.0.13 (rokudev/testing12345)
1) roku-3 10.0.0.100 (admin/demo)
2) roku-2 10.0.0.12 (rokudev/roku4life) current
```

#### Subcommand: `delete|d|del`

Delete a Roku device in your `$HOME/.rokuclirc` file.

```
$ roku-cli devices delete --choice=0
0) stick 192.168.0.13 (rokudev/testing12345) current
1) roku-2 10.0.0.12 (rokudev/roku4life)
```

**Note:** Flags also have shorthand versions too.

- --choice (-c)
- --name (-n)
- --ip (-i)
- --username (-u)
- --password (-p)
- --default (-d)

### Command: `build` (TODO)

This will build a ZIP file to be installed or submitted to Roku for distribution.

```
$ roku-cli build
Building from path: ./src
Build complete: ./build/package.zip

$ roku-cli build ~/some/other/path /var/tmp
Building from path: ~/some/other/path
Build complete: /var/tmp/package.zip
```

### Command: `install` (TODO)

This will build a ZIP file and install it to the currently active Roku device.

```
$ roku-cli install
Building from path: ./src
Build complete: ./build/package.zip
Installing to: roku-2 10.0.0.12

$ roku-cli install ~/some/other/path /var/tmp
Building from path: ~/some/other/path
Build complete: /var/tmp/package.zip
Installing to: roku-2 10.0.0.12
```

### Command: `debug` (TODO)

This will connect to the available telnet ports for debugging information being provided by the Roku while your app is running.

```
$ roku-cli debug
Connected to: roku-2 10.0.0.12
```

## License

Apache 2.0 © [Odd Networks](http://oddnetworks.com)
