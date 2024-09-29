package main

import (
	Shell "clilistener/utils/shell"
	"fmt"
)

func main() {
	shell, path, err := Shell.DetectShell()
	fmt.Println(shell, path)

	if err != nil {
		panic(err)
	}

	history, err := Shell.GetCommandHistory(shell, path)

	if err != nil {
		panic(err)
	}

	topCommands := Shell.GetTopCommands(history, 10)

	fmt.Println(topCommands)
}
