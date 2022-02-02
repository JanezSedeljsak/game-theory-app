package main

import (
	"embed"
	"game-theory-app/core/builder"
)

//go:embed app/public
var fs embed.FS

func main() {
	builder.RunApp(fs)
}
