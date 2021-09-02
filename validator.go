package main

import (
	"fmt"
	"os"
)

// Validator struct defines the structure of errors to be passed in
type Validator struct {
	errors []string
}

func newValidator() *Validator {
	v := new(Validator)

	return v
}

func (v *Validator) errorResponse(message string) {
	msg := message + "\n"
	fmt.Fprintf(os.Stderr, msg)
	os.Exit(1)
}

func (v *Validator) validateInputIsNumber(input []string) {
}

func (v *Validator) flagIsProvided(ports, portRange []string) {
}

func (v *Validator) multipleFlagsNotProvided(ports, portRange []string) {
}

func (v *Validator) rangeOfPorts(rop []string) {
}

func (v *Validator) isValid() error {
	return nil
}
