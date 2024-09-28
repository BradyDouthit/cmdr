package Shell

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// TODO: Add support for separating the args from the command. Should also convert timestamp to a golang time object or number
type Command struct {
	Timestamp string
	Command   string
	Duration  float64
}

func GetCommand(shell, line string) (Command, error) {
	const linePrefix = " : "

	cleanedLine := strings.TrimPrefix(line, linePrefix)

	if !strings.HasSuffix(cleanedLine, "\\") && len(line) > 0 {
		timestamp := sliceBetweenSubstrings(cleanedLine, ":", ":")
		durationStr := sliceBetweenSubstrings(cleanedLine, timestamp+":", ";")
		duration, err := strconv.ParseFloat(durationStr, 32)

		if err != nil {
			return Command{}, fmt.Errorf("could not parse duration: %w", err)
		}

		command := strings.Split(cleanedLine, ";")[1]

		return Command{
			Timestamp: timestamp,
			Command:   command,
			Duration:  duration,
		}, nil
	}

	return Command{}, fmt.Errorf("TODO: Parse multiline commands")
}

func GetCommandHistory(shell, historyFilePath string) ([]Command, error) {
	var history []Command

	data, err := os.ReadFile(historyFilePath)

	if err != nil {
		return history, fmt.Errorf("could not read history file: %w", err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		command, err := GetCommand(shell, line)

		if err != nil {
			continue
		}

		history = append(history, command)
	}

	return history, nil
}

func DetectShell() (string, string, error) {
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

func sliceBetweenSubstrings(str, start, end string) string {
	startIndex := strings.Index(str, start)
	if startIndex == -1 {
		return "" // Start substring not found
	}
	startIndex += len(start)

	endIndex := strings.Index(str[startIndex:], end)
	if endIndex == -1 {
		return "" // End substring not found
	}
	endIndex += startIndex

	return str[startIndex:endIndex]
}
