package main

import (
	"embed"
	"game-theory-app/core/builder"
	"game-theory-app/core/testing"
	"os"
)

//go:embed app/public
var fs embed.FS

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "test" {
		testing.Run()
		return
	}

	builder.RunApp(fs)
}
