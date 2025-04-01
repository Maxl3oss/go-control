package helper

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func KillProcessesByWindowTitle(siteName string) error {
	// Get list of processes using tasklist
	cmd := exec.Command("tasklist", "/v", "/fo", "csv")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get process list: %v", err)
	}

	// Split output into lines
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		// Skip empty lines and header
		if len(line) == 0 || strings.Contains(line, "Image Name") {
			continue
		}

		// Parse CSV (handling quoted fields)
		fields := strings.Split(line, ",")
		if len(fields) < 9 { // tasklist /v gives 9 columns
			continue
		}

		// Last field is Window Title (index 8)
		windowTitle := strings.Trim(fields[8], "\"")
		pidStr := strings.Trim(fields[1], "\"") // PID is second field

		if strings.Contains(windowTitle, siteName) {
			pid, err := strconv.Atoi(pidStr)
			if err != nil {
				continue
			}

			// Kill the process using taskkill
			killCmd := exec.Command("taskkill", "/PID", strconv.Itoa(pid), "/F")
			err = killCmd.Run()
			if err != nil {
				fmt.Printf("Failed to kill process %d: %v\n", pid, err)
			}
		}
	}
	return nil
}

func RunCommandWithStream(c *gin.Context, commandStr string, workingDir string, action string) {
	var shell, flag string
	if runtime.GOOS == "windows" {
		shell, flag = "cmd", "/C"
	} else {
		shell, flag = "sh", "-c"
	}

	command := exec.Command(shell, flag, fmt.Sprintf("cd %s && %s", workingDir, commandStr))
	stdoutPipe, err := command.StdoutPipe()
	if err != nil {
		log.Printf("[%s] Error creating stdout pipe: %v", action, err)
		c.SSEvent("error", fmt.Sprintf("[%s] Failed to initialize output streaming", action))
		c.Writer.Flush()
		return
	}
	stderrPipe, err := command.StderrPipe()
	if err != nil {
		log.Printf("[%s] Error creating stderr pipe: %v", action, err)
		c.SSEvent("error", fmt.Sprintf("[%s] Failed to initialize error streaming", action))
		c.Writer.Flush()
		return
	}

	if err := command.Start(); err != nil {
		log.Printf("[%s] Command failed to start: %v", action, err)
		c.SSEvent("error", fmt.Sprintf("[%s] Failed to start command", action))
		c.Writer.Flush()
		return
	}

	c.SSEvent("status", fmt.Sprintf("[%s] Command started successfully", action))
	c.Writer.Flush()

	done := make(chan bool)
	go func() {
		if _, err := io.Copy(c.Writer, stdoutPipe); err != nil {
			log.Printf("[%s] Error streaming stdout: %v", action, err)
		}
		done <- true
	}()
	go func() {
		if _, err := io.Copy(c.Writer, stderrPipe); err != nil {
			log.Printf("[%s] Error streaming stderr: %v", action, err)
		}
		done <- true
	}()

	<-done
	<-done

	if err := command.Wait(); err != nil {
		log.Printf("[%s] Command execution failed: %v", action, err)
		c.SSEvent("error", fmt.Sprintf("[%s] Command execution failed: %v", action, err))
		c.Writer.Flush()
		return
	}

	c.SSEvent("done", fmt.Sprintf("[%s] Command executed successfully", action))
	c.Writer.Flush()
}

func RunMultiCommandsWithStream(c *gin.Context, commands []string, workingDir string, action string, hiddenCommand string) {
	var shell, flag string
	if runtime.GOOS == "windows" {
		shell, flag = "cmd", "/C"
	} else {
		shell, flag = "sh", "-c"
	}

	for _, commandStr := range commands {
		command := exec.Command(shell, flag, fmt.Sprintf("cd %s && %s", workingDir, commandStr))
		showCommand := FilterCommand(commandStr, hiddenCommand)

		stdoutPipe, err := command.StdoutPipe()
		if err != nil {
			log.Printf("[%s] Error creating stdout pipe: %v", action, err)
			c.SSEvent("error", fmt.Sprintf("[%s] Failed to initialize output streaming for command: %s", action, showCommand))
			c.Writer.Flush()
			return
		}
		stderrPipe, err := command.StderrPipe()
		if err != nil {
			log.Printf("[%s] Error creating stderr pipe: %v", action, err)
			c.SSEvent("error", fmt.Sprintf("[%s] Failed to initialize error streaming for command: %s", action, showCommand))
			c.Writer.Flush()
			return
		}

		if err := command.Start(); err != nil {
			log.Printf("[%s] Command failed to start: %v", action, err)
			c.SSEvent("error", fmt.Sprintf("[%s] Failed to start command: %s", action, showCommand))
			c.Writer.Flush()
			return
		}

		c.SSEvent("status", fmt.Sprintf("[%s] Running command: %s", action, showCommand))
		c.Writer.Flush()

		done := make(chan bool)
		go func() {
			if _, err := io.Copy(c.Writer, stdoutPipe); err != nil {
				log.Printf("[%s] Error streaming stdout for command: %s, %v", action, showCommand, err)
			}
			done <- true
		}()
		go func() {
			if _, err := io.Copy(c.Writer, stderrPipe); err != nil {
				log.Printf("[%s] Error streaming stderr for command: %s, %v", action, showCommand, err)
			}
			done <- true
		}()

		<-done
		<-done

		if err := command.Wait(); err != nil {
			log.Printf("[%s] Command execution failed for command: %s, %v", action, showCommand, err)
			c.SSEvent("error", fmt.Sprintf("[%s] Command execution failed for command: %s, %v", action, showCommand, err))
			c.Writer.Flush()
			return
		}

		c.SSEvent("done", fmt.Sprintf("[%s] Command executed successfully: %s", action, showCommand))
		c.Writer.Flush()
	}

	c.SSEvent("done", fmt.Sprintf("[%s] All commands executed successfully", action))
	c.Writer.Flush()
}

func FilterCommand(commandStr, hiddenCommand string) string {
	// Split the command string into parts
	parts := strings.Split(commandStr, " ")

	// Filter out the hiddenCommand
	var filteredParts []string
	for _, part := range parts {
		if part != hiddenCommand {
			filteredParts = append(filteredParts, part)
		} else {
			filteredParts = append(filteredParts, "********")
		}
	}

	// Join the filtered parts back into a single string
	return strings.Join(filteredParts, " ")
}
