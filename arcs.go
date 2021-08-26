package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

func execWindows(ports, portRange []string) (string, error) {
	var ws syscall.WaitStatus

	for _, port := range ports {
		command := fmt.Sprintf("(Get-NetTCPConnection -LocalPort %s).OwningProcess -Force", port)
		cmd := exec.Command("Stop-Process", "-Id", command)
		if err := cmd.Run(); err != nil {
			if err != nil {
				return "", err
			}

			if exiterr, ok := err.(*exec.ExitError); ok {
				ws = exiterr.Sys().(syscall.WaitStatus)
				return "", exiterr
			}
		}

		ws = cmd.ProcessState.Sys().(syscall.WaitStatus)
		fmt.Printf("Port successfully killed (exit code: %s)\n", []byte(fmt.Sprintf("%d", ws.ExitStatus())))
	}
	return "complete", nil
}

func execUnix(ports, portRange []string) (string, error) {
	var ws syscall.WaitStatus

	for _, port := range ports {
		command := fmt.Sprintf("lsof -i tcp:%s | grep LISTEN | awk '{print $2}' | xargs kill -9", port)
		cmd := exec.Command("bash", "-c", command)

		if err := cmd.Run(); err != nil {
			if err != nil {
				return "", err
			}

			if exiterr, ok := err.(*exec.ExitError); ok {
				ws = exiterr.Sys().(syscall.WaitStatus)
				return "", exiterr
			}
		}

		ws = cmd.ProcessState.Sys().(syscall.WaitStatus)
		fmt.Printf("Port successfully killed (exit code: %s)\n", []byte(fmt.Sprintf("%d", ws.ExitStatus())))
	}

	return "complete", nil
}
