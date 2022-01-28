package main

import (
	"embed"
	"game-theory-app/core"
)

//go:embed app/public
var fs embed.FS

func main() {
	core.RunApp(fs)
}
