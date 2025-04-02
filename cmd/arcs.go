package main

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

func execWindows(ports []string) error {
	var successPorts []string
	var failedPorts []string
	var notFoundPorts []string

	for _, port := range ports {
		// First get the PID of the process using the port
		getPIDCmd := fmt.Sprintf("Get-NetTCPConnection -LocalPort %s -State Listen | Select-Object -ExpandProperty OwningProcess", port)
		cmd := exec.Command("powershell", "-Command", getPIDCmd)

		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("No process found listening on port %s\n", port)
			notFoundPorts = append(notFoundPorts, port)
			continue
		}

		// Clean the output to get just the PID
		pid := strings.TrimSpace(string(output))
		if pid == "" {
			fmt.Printf("No process found listening on port %s\n", port)
			notFoundPorts = append(notFoundPorts, port)
			continue
		}

		// PowerShell might return multiple PIDs (one per line)
		pids := strings.Split(pid, "\n")
		portSuccess := true

		for _, singlePid := range pids {
			singlePid = strings.TrimSpace(singlePid)
			if singlePid == "" {
				continue
			}

			// Now kill the process with the PID
			killCmd := fmt.Sprintf("Stop-Process -Id %s -Force", singlePid)
			killCmdExec := exec.Command("powershell", "-Command", killCmd)

			if err := killCmdExec.Run(); err != nil {
				portSuccess = false
				if exiterr, ok := err.(*exec.ExitError); ok {
					ws := exiterr.Sys().(syscall.WaitStatus)
					fmt.Printf("Failed to kill process (PID: %s) on port %s (exit code: %d)\n",
						singlePid, port, ws.ExitStatus())
				} else {
					fmt.Printf("Error killing process (PID: %s) on port %s: %v\n", singlePid, port, err)
				}
			} else {
				ws := killCmdExec.ProcessState.Sys().(syscall.WaitStatus)
				fmt.Printf("Process (PID: %s) on port %s successfully killed (exit code: %d)\n",
					singlePid, port, ws.ExitStatus())
			}
		}

		if portSuccess {
			successPorts = append(successPorts, port)
		} else {
			failedPorts = append(failedPorts, port)
		}
	}

	// Print summary
	fmt.Println("\n--- Summary ---")
	if len(successPorts) > 0 {
		fmt.Printf("Successfully killed processes on ports: %s\n", strings.Join(successPorts, ", "))
	}
	if len(failedPorts) > 0 {
		fmt.Printf("Failed to kill processes on ports: %s\n", strings.Join(failedPorts, ", "))
	}
	if len(notFoundPorts) > 0 {
		fmt.Printf("No processes found on ports: %s\n", strings.Join(notFoundPorts, ", "))
	}

	// Return error if any port failed
	if len(failedPorts) > 0 {
		return fmt.Errorf("failed to kill processes on some ports")
	}

	return nil
}

func execUnix(ports []string) error {
	var successPorts []string
	var failedPorts []string
	var notFoundPorts []string

	for _, port := range ports {
		// Check if any process is listening on this port
		checkCmd := fmt.Sprintf("lsof -i tcp:%s | grep LISTEN", port)
		checkProcess := exec.Command("bash", "-c", checkCmd)

		checkOutput, err := checkProcess.Output()
		if err != nil || len(checkOutput) == 0 {
			// No process found or command failed
			notFoundPorts = append(notFoundPorts, port)
			continue
		}

		// Get the PID(s) of processes using this port
		pidCmd := fmt.Sprintf("lsof -i tcp:%s | grep LISTEN | awk '{print $2}'", port)
		getPID := exec.Command("bash", "-c", pidCmd)

		pidOutput, err := getPID.Output()
		if err != nil {
			fmt.Printf("Failed to get PID for port %s: %v\n", port, err)
			failedPorts = append(failedPorts, port)
			continue
		}

		// Split the output in case there are multiple PIDs
		pids := strings.Split(strings.TrimSpace(string(pidOutput)), "\n")

		portSuccess := true
		// Kill each process
		for _, pid := range pids {
			if pid == "" {
				continue
			}

			killCmd := fmt.Sprintf("kill -9 %s", pid)
			killProcess := exec.Command("bash", "-c", killCmd)

			if err := killProcess.Run(); err != nil {
				portSuccess = false
				if exiterr, ok := err.(*exec.ExitError); ok {
					ws := exiterr.Sys().(syscall.WaitStatus)
					fmt.Printf("Failed to kill process (PID: %s) on port %s (exit code: %d)\n",
						pid, port, ws.ExitStatus())
				} else {
					fmt.Printf("Error killing process (PID: %s) on port %s: %v\n", pid, port, err)
				}
			} else {
				fmt.Printf("Process (PID: %s) on port %s successfully killed\n", pid, port)
			}
		}

		if portSuccess {
			successPorts = append(successPorts, port)
		} else {
			failedPorts = append(failedPorts, port)
		}
	}

	// Print summary
	fmt.Println("\n--- Summary ---")
	if len(successPorts) > 0 {
		fmt.Printf("Successfully killed processes on ports: %s\n", strings.Join(successPorts, ", "))
	}
	if len(failedPorts) > 0 {
		fmt.Printf("Failed to kill processes on ports: %s\n", strings.Join(failedPorts, ", "))
	}
	if len(notFoundPorts) > 0 {
		fmt.Printf("No processes found on ports: %s\n", strings.Join(notFoundPorts, ", "))
	}

	// Return error if any port failed
	if len(failedPorts) > 0 {
		return fmt.Errorf("failed to kill processes on some ports")
	}

	return nil
}
