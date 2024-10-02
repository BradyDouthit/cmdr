package ui

import (
	Shell "cmdr/utils/shell"
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func RenderTopCommands(command []Shell.CommandCount) {
	for _, command := range command {
		primaryStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F0F0F0"))

		commandStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#61AFEF")).
			Background(lipgloss.Color("#011b30")).
			Bold(true)

		countStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F0F0F0")).
			Bold(true)

		output := primaryStyle.Render("You have used ") +
			commandStyle.Render(command.Command) +
			primaryStyle.Render(" ") +
			countStyle.Render(fmt.Sprintf("%d", command.Count)) +
			primaryStyle.Render(" times")

		fmt.Println(output)
	}
}

func RenderMistakes(commands []Shell.CommandCount) {
	for _, command := range commands {
		primaryStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F0F0F0"))

		commandStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B6B")).
			Background(lipgloss.Color("#4A0E0E")).
			Bold(true)

		countStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F0F0F0")).
			Bold(true)

		output := primaryStyle.Render("You have used ") +
			commandStyle.Render(command.Command) +
			primaryStyle.Render(" ") +
			countStyle.Render(fmt.Sprintf("%d", command.Count)) +
			primaryStyle.Render(" times but it does not exist")

		fmt.Println(output)
	}
}

func RenderTime(elapsed time.Duration) {
	subtleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#808080"))
	fmt.Println(subtleStyle.Render(fmt.Sprintf("Execution time: %v", elapsed)))
}
