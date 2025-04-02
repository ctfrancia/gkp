package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/ctfrancia/gkp/cmd/cli"
)

func main() {
	ports := cli.PortsToKill()
	xps := strings.Split(ports.Ports, " ")
	xrps := strings.Split(ports.PortRange, " ")

	if len(xps) != 0 {
		// check if we have ports

		killPorts(xps)
	} else if len(xrps) != 0 {
		// the user wants to kill a range of ports

		killPortsRange(xrps)
	} else {
		fmt.Println("an incorrect flag was provided or no flags were provided")
		os.Exit(1)
	}
}

// killPorts will kill a list of ports
func killPorts(ports []string) {
	var err error

	switch runtime.GOOS {
	case "windows":
		err = execWindows(ports)
	case "darwin", "linux", "freebsd", "openbsd", "netbsd":
		// These are all Unix-like systems that should work with execUnix
		err = execUnix(ports)
	default:
		fmt.Printf("Unsupported operating system: %s\n", runtime.GOOS)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error killing ports: %v\n", err)
		os.Exit(1)
	}
}

// killPortsRange will kill a range of ports
// TODO: implement this
func killPortsRange(portsRange []string) {
	if runtime.GOOS == "windows" {
	} else {
	}
}
