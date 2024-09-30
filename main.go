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

	UI.RenderTopCommands(topCommands)

	os.Exit(0)
}
