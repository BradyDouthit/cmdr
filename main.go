package main

import (
	ui "github.com/gizak/termui/v3"
	// "strconv"
	"trendify/utils/shell"
	UIHelpers "trendify/utils/uiHelpers"
)

// TODO: Add different CLI flags for charts, suggestions, etc
func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	shell, path, err := Shell.DetectShell()

	if err != nil {
		panic(err)
	}

	history, err := Shell.GetCommandHistory(shell, path)

	if err != nil {
		panic(err)
	}

	chart := UIHelpers.BuildBarChart(history[:5])

	ui.Render(chart)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
