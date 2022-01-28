package core

import "sync"

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

type tictactoe struct {
	sync.Mutex
	board [][]int8
}

type GameStatus struct {
	Board  [][]int8 `json:"board"`
	Winner int      `json:"winner"`
	IsDone bool     `json:"isdone"`
}
