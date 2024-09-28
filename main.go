package main

import (
	"clilistener/utils/shell"
	"fmt"
	"strconv"
)

func main() {
	shell, path, err := Shell.DetectShell()

	if err != nil {
		panic(err)
	}

	history, err := Shell.GetCommandHistory(shell, path)

	if err != nil {
		panic(err)
	}

	fmt.Println(strconv.Itoa(len(history)) + " commands found")
}
