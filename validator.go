package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Validator struct defines the structure of errors to be passed in
type Validator struct {
	errors []string
}

func newValidator() *Validator {
	v := new(Validator)

	return v
}

func (v *Validator) addError(e string) {
	v.errors = append(v.errors, e)
}

func (v *Validator) errorResponse(message []string) {
	msg := strings.Join(message, ", ")
	fmt.Println(msg)
	os.Exit(1)
}

func (v *Validator) multipleFlagsNotProvided(ports, portRange []string) {
	if len(ports) > 1 && len(portRange) > 1 {
		v.addError("cannot have range and port flag at the same time")
	}
}

func (v *Validator) rangeOfPortsAreNumbers(ports []string) []int {
	var numPorts []int
	for _, port := range ports {
		if _, err := strconv.Atoi(port); err != nil {
			msg := fmt.Sprintf("Error: port argument is not a number: %s \n", port)
			v.addError(msg)
		}

		p, _ := strconv.Atoi(port)
		numPorts = append(numPorts, p)
	}

	return numPorts
}

func (v *Validator) rangeAreAscending(p1, p2 int) {
	if p1 > p2 {
		v.addError("first argument must be less than second argument")
	}
}

func (v *Validator) oneFlagProvided(ports, rangeOfPorts []string) {
	if len(ports) < 1 && len(rangeOfPorts) < 1 {
		v.errors = append(v.errors, "no flags provided\n")
	}
}

func (v *Validator) isValid() {
	if len(v.errors) > 0 {
		v.errorResponse(v.errors)
	}
}
