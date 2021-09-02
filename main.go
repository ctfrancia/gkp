package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/ctfrancia/gkp/cmd/cli"
)

func main() {
	ports := cli.PortsToKill()
	v := newValidator()

	var xps []string = strings.Split(ports.Ports, " ")
	var xrps []string = strings.Split(ports.PortRange, " ")

	v.validateInputIsNumber(xps)
	v.validateInputIsNumber(xrps)

	if len(xps) < 1 && len(xrps) < 1 {
		fmt.Fprintf(os.Stderr, "no flags provided\n")
		os.Exit(1)
	}

	if len(xps) > 1 && len(xrps) > 1 {
		fmt.Println("cannot have range and port flag at the same time")
		os.Exit(1)
	}

	// checks if each port provided is an integer
	for _, p := range xps {
		if _, err := strconv.Atoi(p); err != nil {
			fmt.Fprintf(os.Stderr, "Error: port argument is not a number %s \n", p)
			os.Exit(1)
		}
	}

	if len(xrps) > 1 {
		fmt.Println("shoiuld nto be here", xrps)
		// check if the two numbers provided are valid integers
		for _, rp := range xrps {
			if _, err := strconv.Atoi(rp); err != nil {
				fmt.Fprintf(os.Stderr, "Error: port argument is not a number %s \n", rp)
				os.Exit(1)
			}
		}
	}

	if len(xrps) != 2 {
		fmt.Fprintf(os.Stderr, "only two numbers are accepatble for a range, eg: '3000 3002'")
		os.Exit(1)
	}

	// checks to make sure that left to right the numbers increase
	if xrps[0] > xrps[1] {
		fmt.Fprintf(os.Stderr, "first argument must be less than second argument")
		os.Exit(1)
	}

	if runtime.GOOS == "windows" {
		msg, err := execWindows(xps, xrps)
		if err != nil {
			// fmt.Fprintf(os.Stderr, err.Error())
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
