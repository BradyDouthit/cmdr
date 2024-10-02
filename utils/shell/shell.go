package Shell

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"sort"
	"strings"
)

type Command struct {
	Command string   // The command itself
	Args    []string // The arguments passed to the command
	Valid   bool
}

type CommandCount struct {
	Command string
	Count   int
	Valid   bool
}

var BuiltinCommands = []string{"alias", "bg",
	"bind", "break", "builtin", "case", "cd", "command", "compgen", "complete", "continue",
	"declare", "dirs", "disown", "echo", "enable", "eval", "exec", "exit", "export", "fc",
	"fg", "getopts", "hash", "help", "history", "if", "jobs", "kill", "let", "local", "logout",
	"popd", "printf", "pushd", "pwd", "read", "readonly", "return", "set", "shift", "shopt",
	"source", "suspend", "test", "times", "trap", "type", "typeset", "ulimit", "umask",
	"un‐alias", "unset", "until", "wait", "while"}

func GetCommand(shell, line string) (Command, error) {
	if shell == "bash" {
		command, err := parseCommandOnly(line)
		return command, err
	} else if shell == "zsh" {
		command, err := getZshCommand(line)
		return command, err
	}

	// TODO: Add support for other shells
	return Command{}, nil
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
	var shellSuffix string

	if runtime.GOOS == "windows" {
		executableName := shell[strings.LastIndex(shell, "\\")+1:]
		shellSuffix = strings.Split(executableName, ".")[0]
	} else {
		shellSuffix = shell[strings.LastIndex(shell, "/")+1:]
	}

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

func GetUniqueCommandCounts(history []Command, count int, includeArgs bool) []CommandCount {
	commandCounts := make(map[string]CommandCount)

	for _, cmd := range history {
		if includeArgs {
			fullCommand := cmd.Command + " " + strings.Join(cmd.Args, " ")
			prevCount := commandCounts[fullCommand].Count
			commandCounts[fullCommand] = CommandCount{Command: fullCommand, Count: prevCount + 1, Valid: cmd.Valid}
		} else {
			prevCount := commandCounts[cmd.Command].Count
			commandCounts[cmd.Command] = CommandCount{Command: cmd.Command, Count: prevCount + 1, Valid: cmd.Valid}
		}
	}

	var topCommands []CommandCount
	// Log the counts for each of the top N commands
	for _, count := range commandCounts {
		topCommands = append(topCommands, count)
	}

	sort.Slice(topCommands, func(i, j int) bool {
		return topCommands[i].Count > topCommands[j].Count
	})

	if len(topCommands) > count {
		topCommands = topCommands[:count]
	}

	return topCommands
}

func GetCommandExists(command string) bool {
	// Validate that this works on windows
	if slices.Contains(BuiltinCommands, command) {
		return true
	}

	_, err := exec.LookPath(command)

	if err == nil {
		return true
	}

	// works on my machine lol
	// Really though, this isn't a good idea.
	if runtime.GOOS == "windows" {
		commandErr := exec.Command("bash", "-c", fmt.Sprintf("command -v %s", command)).Run()

		return commandErr == nil
	}

	return false
}

func parseCommandOnly(line string) (Command, error) {
	parts := strings.Split(line, " ")

	if len(parts) == 0 {
		return Command{}, errors.New("Found an empty line in the history file")
	}

	mainCommand := strings.TrimSpace(parts[0])

	if len(mainCommand) == 0 {
		return Command{}, errors.New("Found an empty command in the history file")
	}

	return Command{
		Command: mainCommand,
		Args:    parts[1:],
		Valid:   GetCommandExists(mainCommand),
	}, nil

}

func getZshCommand(line string) (Command, error) {
	rawCommand := sliceBetweenSubstrings(line, ";", "")
	cmd, err := parseCommandOnly(rawCommand)

	return cmd, err
}

func sliceBetweenSubstrings(str, start, end string) string {
	startIndex := strings.Index(str, start)
	if startIndex == -1 {
		return "" // Start substring not found
	}
	startIndex += len(start)

	endIndex := strings.Index(str[startIndex:], end)

	// If the end substring is empty, return the rest of the string
	if endIndex == -1 || end == "" {
		return str[startIndex:] // End substring not found
	}

	endIndex += startIndex

	return str[startIndex:endIndex]
}
