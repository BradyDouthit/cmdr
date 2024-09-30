package ui

import (
	Shell "clilistener/utils/shell"
	"fmt"

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
