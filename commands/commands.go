package commands

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

// Run a command with stdout printed to console
func Run(command string) error {
	return runCommand(command, true)
}

// Run a command without stdout printed
func RunSilent(command string) error {
	return runCommand(command, false)
}

// Run a command and return output
func Output(command string) (string, error) {
	cmd, err := getExecCommand(command)
	if err != nil {
		return "", err
	}

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func runCommand(command string, displayOutput bool) error {
	cmd, err := getExecCommand(command)
	if err != nil {
		return err
	}

	if displayOutput {
		cmd.Stdout = os.Stdout
	}

	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func getExecCommand(command string) (*exec.Cmd, error) {
	command = normalizeCommand(command)
	if len(command) == 0 {
		return nil, errors.New("no command provided")
	}

	args := strings.Split(command, " ")
	return exec.Command(args[0], args[1:]...), nil
}

func normalizeCommand(command string) string {
	command = strings.TrimSpace(command)
	for strings.Contains(command, "  ") {
		command = strings.ReplaceAll(command, "  ", " ")
	}
	return command
}
