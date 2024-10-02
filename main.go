package main

import (
	Shell "cmdr/utils/shell"
	UI "cmdr/utils/ui"
	"flag"
	"os"
	"time"
)

func exit(code int, startTime time.Time) {
	elapsed := time.Since(startTime)
	UI.RenderTime(elapsed)
	os.Exit(code)
}

func main() {
	mainStart := time.Now()
	includeArgsShort := flag.Bool("A", false, "Include arguments in the output")
	includeArgsLong := flag.Bool("args", false, "Include arguments in the output")

	showMistakesLong := flag.Bool("mistakes", false, "Show mistakes (commands that aren't available on your system) in the output")
	showMistakesShort := flag.Bool("M", false, "Show mistakes (commands that aren't available on your system) in the output")

	showValidLong := flag.Bool("valid", false, "Show valid commands in the output")
	showValidShort := flag.Bool("V", false, "Show valid commands in the output")

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

	if *showValidShort || *showValidLong {
		uniqueCommands := Shell.GetUniqueCommandCounts(history, 10000, *includeArgsShort || *includeArgsLong)
		failedCommands := Shell.GetCommandsExist(uniqueCommands, *topN, true)
		UI.RenderValid(failedCommands)

		exit(0, mainStart)
	}

	if *showMistakesLong || *showMistakesShort {
		uniqueCommands := Shell.GetUniqueCommandCounts(history, 10000, *includeArgsShort || *includeArgsLong)
		failedCommands := Shell.GetCommandsExist(uniqueCommands, *topN, false)
		UI.RenderMistakes(failedCommands)

		exit(0, mainStart)
	}

	topCommands := Shell.GetUniqueCommandCounts(history, *topN, *includeArgsShort || *includeArgsLong)
	UI.RenderTopCommands(topCommands)

	exit(0, mainStart)
}
