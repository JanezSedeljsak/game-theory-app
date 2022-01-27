package main

import (
	"fmt"
	"os"
	"tictactoe-minmax/core"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "cli" {
		core.RunConsoleGame()
	} else {
		fmt.Println("Run main script")
	}
}
