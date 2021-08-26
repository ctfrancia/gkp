package main

import (
	"errors"
	"fmt"
	"os/exec"
)

func execWindows(ports, portRange []string) (string, error) {
	if len(ports) > 1 && len(portRange) > 1 {
		return "", errors.New("cannot have range and port flag at the same time")
	}

	for _, port := range ports {
		command := fmt.Sprintf("(Get-NetTCPConnection -LocalPort %s).OwningProcess -Force", port)
		exec.Command("Stop-Process", "-Id", command)
	}

	return "", nil
}

func execUnix(ports, portRange []string) (string, error) {
	if len(ports) > 1 && len(portRange) > 1 {
		return "", errors.New("cannot have range and port flag at the same time")
	}

	for _, port := range ports {
		command := fmt.Sprintf("lsof -i tcp:%s | grep LISTEN | awk '{print $2}' | xargs kill -9", port)
		exec.Command("bash", "-c", command)
	}

	return "", nil
}
