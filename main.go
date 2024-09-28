package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	shell, path := detectShell()

	fmt.Println("Found shell and path:", shell, path)
}

func detectShell() (string, string) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic("Could not get home directory")
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

	return shellSuffix, historyFilePath
}
