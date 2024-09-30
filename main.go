package main

import (
	Shell "clilistener/utils/shell"
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

func renderTopCommands(command Shell.CommandCount) {
	primaryStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F0F0F0"))

	commandStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#61AFEF")).
		Background(lipgloss.Color("#011b30")).
		Bold(true)

	countStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#98C379")).
		Bold(true)

	output := primaryStyle.Render("You have used ") +
		commandStyle.Render(command.Command) +
		primaryStyle.Render(" ") +
		countStyle.Render(fmt.Sprintf("%d", command.Count)) +
		primaryStyle.Render(" times")

	fmt.Println(output)
}

func main() {
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
		renderTopCommands(command)
	}

	os.Exit(0)
}
