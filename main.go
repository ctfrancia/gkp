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
	v := newValidator()
	var intRangePorts []int

	var xps []string = strings.Split(ports.Ports, " ")
	var xrps []string = strings.Split(ports.PortRange, " ")

	_ = v.rangeOfPortsAreNumbers(xps)
	if len(xrps) == 2 {
		intRangePorts = v.rangeOfPortsAreNumbers(xrps)
	}

	v.oneFlagProvided(xps, xrps)
	v.multipleFlagsNotProvided(xps, xrps)

	if len(xrps) > 1 {
		v.rangeAreAscending(intRangePorts[0], intRangePorts[1])
	}

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
