package main

import (
	"clilistener/utils/shell"
	"fmt"
)

func main() {
	shell, path, err := Shell.DetectShell()

	if err != nil {
		panic(err)
	}

	fmt.Println("Found shell and path:", shell, path)

	history, err := Shell.GetCommandHistory(shell, path)

	if err != nil {
		panic(err)
	}

	fmt.Println("History:", len(history))
}
