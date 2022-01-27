package main

import (
	"embed"
	"os"
	"tictactoe-minmax/core"
)

//go:embed app/public
var fs embed.FS

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "cli" {
		core.RunConsoleGame()
	} else {
		core.BuildUserInterface(fs)
	}
}
