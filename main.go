package main

import (
	Shell "clilistener/utils/shell"
	"flag"
	"fmt"
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

	for _, command := range topCommands {
		fmt.Printf("You have used %s %d times\n", command.Command, command.Count)
	}

	os.Exit(0)
}
