package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TODO: Add support for separating the args from the command. Should also convert timestamp to a golang time object or number
type Command struct {
	Timestamp string
	Command   string
}

func main() {
	shell, path, err := detectShell()

	if err != nil {
		panic(err)
	}

	fmt.Println("Found shell and path:", shell, path)

	history, err := getCommandHistory(shell, path)

	if err != nil {
		panic(err)
	}

	fmt.Println("History:", len(history))
}

func getCommand(shell, line string) (*Command, error) {
	return nil, nil
}

func getCommandHistory(shell, historyFilePath string) ([]string, error) {
	var history []string

	data, err := os.ReadFile(historyFilePath)

	if err != nil {
		return history, fmt.Errorf("could not read history file: %w", err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		_, err := getCommand(shell, line)

		if err != nil {
			return history, fmt.Errorf("could not parse command: %w", err)
		}

		fmt.Println("Line:", line)
	}

	return history, nil
}

func detectShell() (string, string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", "", fmt.Errorf("could not get home directory: %w", err)
	}

	shell := os.Getenv("SHELL")
	shellSuffix := shell[strings.LastIndex(shell, "/")+1:]

	// The path that the CLI history is saved to
	var historyFilePath string

	switch shellSuffix {
	case "bash":
		historyFilePath = filepath.Join(homeDir, ".bash_history")
		break
	case "zsh":
		historyFilePath = filepath.Join(homeDir, ".zsh_history")
		break
	case "fish":
		historyFilePath = filepath.Join(homeDir, ".local", "share", "fish", "fish_history")
		break
	default:
		historyFilePath = filepath.Join(homeDir, ".bash_history")
		break
	}

	if _, err := os.Stat(historyFilePath); errors.Is(err, os.ErrNotExist) {
		return shellSuffix, "", fmt.Errorf("could not find history file at %s", historyFilePath)
	}

	return shellSuffix, historyFilePath, nil
}
