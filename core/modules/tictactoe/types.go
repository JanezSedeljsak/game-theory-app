package tictactoe

import (
	"game-theory-app/core/globals"
	"sync"
)

const Size = 3

type State struct {
	sync.Mutex
	board [][]int8
}

type GameStatus struct {
	Board  [][]int8        `json:"board"`
	Winner int             `json:"winner"`
	IsDone bool            `json:"isdone"`
	Coords []globals.Coord `json:"coords"`
}
