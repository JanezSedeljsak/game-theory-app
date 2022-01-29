package core

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"game-theory-app/core/tictactoe"

	"github.com/zserge/lorca"
)

func RunApp(fs embed.FS) {
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}

	ui, err := lorca.New("", "", 1024, 720, args...)
	if err != nil {
		log.Fatal(err)
	}

	ui.Bind("start", func() {
		log.Println("UI is ready")
	})

	ttt := &tictactoe.State{}
	ui.Bind("mutate", ttt.Mutate)
	ui.Bind("init", ttt.Init)
	ui.Bind("status", ttt.Status)
	ui.Bind("random", ttt.RandomMove)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	go http.Serve(ln, http.FileServer(http.FS(fs)))
	ui.Load(fmt.Sprintf("http://%s/app/public", ln.Addr()))
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	defer ui.Close()
	defer ln.Close()
	log.Println("exiting...")
}
