package main

import (
	Shell "cmdr/utils/shell"
	UI "cmdr/utils/ui"
	"os"
	"strconv"
	"time"

	"github.com/BradyDouthit/switchboard"
)

func exit(code int, startTime time.Time) {
	elapsed := time.Since(startTime)
	UI.RenderTime(elapsed)
	os.Exit(code)
}

func addModifierArgs(cmd *switchboard.Command, topN *int, includeArgs *bool) {
	cmd.Flag("t", "top", "Filter for the top N commands", false, func(s string) error {
		t, err := strconv.Atoi(s)

		if err != nil {
			return err
		}

		*topN = t
		return nil
	})

	cmd.BoolFlag("a", "args", "Include arguments in the resulting commands", func(b bool) error {
		*includeArgs = b
		return nil
	})
}

func main() {
	mainStart := time.Now()
	shell, path, config, err := Shell.DetectShell()

	if err != nil {
		panic(err)
	}

	aliases, err := Shell.GetAliases(config)

	if err != nil {
		panic(err)
	}

	history, err := Shell.GetCommandHistory(shell, path, aliases)

	if err != nil {
		panic(err)
	}

	app := switchboard.New()

	app.Command("valid", "Filter for VALID commands you have run", func(c *switchboard.Command) {
		var topN int = 5
		var includeArgs bool = false
		var validCommands []Shell.CommandCount

		// Updates topN and includeArgs
		addModifierArgs(c, &topN, &includeArgs)

		c.Run(func() {
			uniqueCommands := Shell.GetUniqueCommandCounts(history, 10000, includeArgs)
			for _, command := range uniqueCommands {
				if command.Valid {
					validCommands = append(validCommands, command)
				}
			}

			if len(validCommands) > topN {
				val := validCommands[:topN]
				UI.RenderValid(val, aliases)
			} else {
				UI.RenderValid(validCommands, aliases)
			}
		})
	})

	app.Command("invalid", "Filter for INVALID commands you have run", func(c *switchboard.Command) {
		var topN int = 5
		var includeArgs bool = false

		// Updates topN and includeArgs
		addModifierArgs(c, &topN, &includeArgs)

		c.Run(func() {
			var invalidCommands []Shell.CommandCount
			uniqueCommands := Shell.GetUniqueCommandCounts(history, 10000, includeArgs)

			for _, command := range uniqueCommands {
				if !command.Valid {
					invalidCommands = append(invalidCommands, command)
				}
			}

			if len(invalidCommands) > topN {
				inv := invalidCommands[:topN]
				UI.RenderInvalid(inv)
			} else {
				UI.RenderInvalid(invalidCommands)
			}
		})
	})

	app.Run()

	exit(0, mainStart)
}
