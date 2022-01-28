package main

import (
	"embed"
	"game-theory-challenge/core"
)

//go:embed app/public
var fs embed.FS

func main() {
	core.RunApp(fs)
}
