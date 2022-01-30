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

func RegisterActions(args []string, ln net.Listener) lorca.UI {
	ui, err := lorca.New("", "", 960, 740, args...)
	if err != nil {
		log.Fatal(err)
	}

	ttt := &tictactoe.Actions{}
	ui.Bind("ttt_init", ttt.Init)
	ui.Bind("ttt_mutateAI", ttt.Mutate)
	ui.Bind("ttt_mutateRand", ttt.RandomMove)
	ui.Bind("ttt_multiplayer", ttt.Multiplayer)

	cf := &connect4.Actions{}
	ui.Bind("cf_init", cf.Init)
	ui.Bind("cf_mutateAI", cf.RandomMove) // not yet implemented -> use random for now
	ui.Bind("cf_mutateRand", cf.RandomMove)
	ui.Bind("cf_multiplayer", cf.Multiplayer)

	ui.Load(fmt.Sprintf("http://%s/app/public", ln.Addr()))
	return ui
}

func BuildConnection(fs embed.FS) net.Listener {
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

	ln := BuildConnection(fs)
	ui := RegisterActions(args, ln)

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	defer ui.Close()
	defer ln.Close()
}
