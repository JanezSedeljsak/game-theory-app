package tictactoe

import "sync"

const Size = 3

type Coord struct {
	Row int
	Col int
}

type Response struct {
	Coords Coord
	Value  int
}

type State struct {
	sync.Mutex
	board [][]int8
}

type GameStatus struct {
	Board  [][]int8 `json:"board"`
	Winner int      `json:"winner"`
	IsDone bool     `json:"isdone"`
	Coords []Coord  `json:"coords"`
}
