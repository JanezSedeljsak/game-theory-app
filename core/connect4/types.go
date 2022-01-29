package connect4

import "sync"

const Height int = 6
const Width int = 7

type Coord struct {
	Row int
	Col int
}

type Response struct {
	Coords Coord
	Value  int
}

type dp struct {
	Memo map[int]Response
}

type State struct {
	sync.Mutex
	board [][]bool
}

type GameStatus struct {
	Board  [][]bool `json:"board"`
	Winner int      `json:"winner"`
	IsDone bool     `json:"isdone"`
	Coords []Coord  `json:"coords"`
}
