package main

import (
	"fmt"
	"os/exec"
)

func execWindows(ports, portRange []string) (string, error) {
	command := fmt.Sprintf("(Get-NetTCPConnection -LocalPort %s).OwningProcess -Force", port)
	exec.Command("Stop-Process", "-Id", command)

	return "", nil
}

func execUnix(ports, portRange []string) (string, error) {
	command := fmt.Sprintf("lsof -i tcp:%s | grep LISTEN | awk '{print $2}' | xargs kill -9", port)
	exec.Command("bash", "-c", command)

	return "", nil
}
