package connect4

import "sync"

const Height int = 6
const Width int = 7

type State struct {
	sync.Mutex
	board Board
}
