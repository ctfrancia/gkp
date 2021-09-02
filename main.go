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
	var intRangePorts []int
	v := newValidator()

	var xps []string = strings.Split(ports.Ports, " ")
	var xrps []string = strings.Split(ports.PortRange, " ")

	_ = v.rangeOfPortsAreNumbers(xps)
	// checks if the first item is the defaulted "" value
	if xrps[0] != "" {
		intRangePorts = v.rangeOfPortsAreNumbers(xrps)
		v.rangeAreAscending(intRangePorts[0], intRangePorts[1])
	}

	v.oneFlagProvided(xps, xrps)
	v.multipleFlagsNotProvided(xps, xrps)

	v.isValid()

	if runtime.GOOS == "windows" {
		msg, err := execWindows(xps, xrps)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(msg)
		os.Exit(0)
	} else {
		msg, err := execUnix(xps, xrps)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(msg)
		os.Exit(1)
	}
}
