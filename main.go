package main

import (
	Shell "clilistener/utils/shell"
	UI "clilistener/utils/ui"
	"flag"
	"os"
)

func main() {
	includeArgsShort := flag.Bool("A", false, "Include arguments in the output")
	includeArgsLong := flag.Bool("args", false, "Include arguments in the output")
	showMistakesLong := flag.Bool("mistakes", false, "Show mistakes (commands that don't exist in the PATH) in the output")
	showMistakesShort := flag.Bool("M", false, "Show mistakes (commands that don't exist in the PATH) in the output")
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

	if *showMistakesLong || *showMistakesShort {
		uniqueCommands := Shell.GetUniqueCommandCounts(history, 999, *includeArgsShort || *includeArgsLong)
		failedCommands := Shell.GetFailedCommands(uniqueCommands, *topN)
		UI.RenderMistakes(failedCommands)
		os.Exit(0)
	}

	topCommands := Shell.GetUniqueCommandCounts(history, *topN, *includeArgsShort || *includeArgsLong)
	UI.RenderTopCommands(topCommands)

	os.Exit(0)
}
