package Shell

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
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
		command := ""

		if err != nil {
			return Command{Timestamp: timestamp, Command: command, Duration: 0}, fmt.Errorf("could not parse duration: %w", err)
		}

		return Command{
			Timestamp: timestamp,
			Command:   command,
			Duration:  duration,
		}, nil
	} else if !strings.HasPrefix(cleanedLine, ":") {
		// TODO: Handle no timestamp
		return Command{
			Timestamp: "",
			Command:   cleanedLine,
			Duration:  0,
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
		command, _ := GetCommand(shell, line)

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

func GetTopCommands(history []Command, count int) []Command {
	topCommands := make(map[string]int)

	fmt.Println(topCommands)
	for _, command := range history {
		topCommands[command.Command]++
	}

	// Create a slice of Command structs to store the top commands
	var result []Command

	// Convert map to slice of key-value pairs
	pairs := make([]struct {
		cmd   string
		count int
	}, 0, len(topCommands))
	for cmd, count := range topCommands {
		pairs = append(pairs, struct {
			cmd   string
			count int
		}{cmd, count})
	}

	// Sort the pairs slice by count in descending order
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	// Take the top N commands
	for i := 0; i < count && i < len(pairs); i++ {
		result = append(result, Command{Command: pairs[i].cmd})
	}

	return result
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
