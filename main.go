package main

import (
	Shell "clilistener/utils/shell"
	"flag"
	"fmt"
	"os"
)

func main() {
	includeArgs := flag.Bool("A", false, "Include arguments in the output")
	includeArgsLong := flag.Bool("args", false, "Include arguments in the output")
	flag.Parse()
	shell, path, err := Shell.DetectShell()

	if err != nil {
		panic(err)
	}

	history, err := Shell.GetCommandHistory(shell, path)

	if err != nil {
		panic(err)
	}

	if *includeArgs || *includeArgsLong {
		topCommands := Shell.GetTopCommands(history, 5)

		for _, command := range topCommands {
			fmt.Printf("You have used %s %d times\n", command.Command, command.Count)
		}

		os.Exit(0)
	}
}
