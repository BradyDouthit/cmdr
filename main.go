package main

import (
	Shell "clilistener/utils/shell"
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	// Primary Text Style
	primaryStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F0F0F0")).
		Background(lipgloss.Color("#1E1E1E"))

	// Command Style
	commandStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#61AFEF")).
		Bold(true)

	// Count Style
	countStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#98C379")).
		Bold(true)

	includeArgsShort := flag.Bool("A", false, "Include arguments in the output")
	includeArgsLong := flag.Bool("args", false, "Include arguments in the output")
	topN := flag.Int("top", 5, "Number of top commands to display")
	flag.Parse()
	shell, path, err := Shell.DetectShell()

	if err != nil {
		panic(err)
	}

	history, err := Shell.GetCommandHistory(shell, path)

	if err != nil {
		panic(err)
	}

	topCommands := Shell.GetTopCommands(history, *topN, *includeArgsShort || *includeArgsLong)

	for _, command := range topCommands {
		output := primaryStyle.Render("You have used ") +
			commandStyle.Render(command.Command) +
			primaryStyle.Render(" ") +
			countStyle.Render(fmt.Sprintf("%d", command.Count)) +
			primaryStyle.Render(" times")

		fmt.Println(output)
	}

	os.Exit(0)
}
