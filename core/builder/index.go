package builder

import (
	"embed"
	"fmt"
	"game-theory-app/core/modules/connect4"
	"game-theory-app/core/modules/tictactoe"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"github.com/zserge/lorca"
)

func BuildServices(args []string, ln net.Listener) lorca.UI {
	ui, err := lorca.New("", "", 1024, 720, args...)
	if err != nil {
		log.Fatal(err)
	}

	ttt := &tictactoe.State{}
	ui.Bind("ttt_mutateAI", ttt.Mutate)
	ui.Bind("ttt_init", ttt.Init)
	ui.Bind("ttt_multiplayer", ttt.Multiplayer)
	ui.Bind("ttt_mutateRand", ttt.RandomMove)

	cf := &connect4.State{}
	ui.Bind("cf_init", cf.Init)

	ui.Load(fmt.Sprintf("http://%s/app/public", ln.Addr()))
	return ui
}

func BuildServer(fs embed.FS) net.Listener {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	go http.Serve(ln, http.FileServer(http.FS(fs)))
	return ln
}

func RunApp(fs embed.FS) {
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}

	ln := BuildServer(fs)
	ui := BuildServices(args, ln)

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	defer ui.Close()
	defer ln.Close()
}
