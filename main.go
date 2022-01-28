package main

import (
	"embed"
	"tictactoe-minmax/core"
)

//go:embed app/public
var fs embed.FS

func main() {
	core.RunApp(fs)
}
