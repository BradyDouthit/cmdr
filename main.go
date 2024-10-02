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

	showInvalidLong := flag.Bool("invalid", false, "Show invalid (commands that aren't available on your system) in the output")
	showInvalidShort := flag.Bool("I", false, "Show invalid (commands that aren't available on your system) in the output")

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

		var validCommands []Shell.CommandCount

		for _, command := range uniqueCommands {
			if command.Valid {
				validCommands = append(validCommands, command)
			}
		}

		if len(validCommands) > *topN {
			val := validCommands[:*topN]
			UI.RenderValid(val)
		} else {
			UI.RenderValid(validCommands)
		}

		exit(0, mainStart)
	}

	if *showInvalidLong || *showInvalidShort {
		uniqueCommands := Shell.GetUniqueCommandCounts(history, 10000, *includeArgsShort || *includeArgsLong)

		var invalidCommands []Shell.CommandCount

		for _, command := range uniqueCommands {
			if !command.Valid {
				invalidCommands = append(invalidCommands, command)
			}
		}

		if len(invalidCommands) > *topN {
			inv := invalidCommands[:*topN]
			UI.RenderInvalid(inv)
		} else {
			UI.RenderInvalid(invalidCommands)
		}

		exit(0, mainStart)
	}

	topCommands := Shell.GetUniqueCommandCounts(history, *topN, *includeArgsShort || *includeArgsLong)
	UI.RenderTopCommands(topCommands)

	exit(0, mainStart)
}
