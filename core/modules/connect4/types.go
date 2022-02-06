package connect4

import (
	"encoding/json"
	"log"
	"sync"
)

const Height int = 6
const Width int = 7

type Actions struct {
	sync.Mutex
	board   Board
	History Stack
}

type dp struct {
	Memo     map[uint64]int8
	MaxDepth int8
}

type Coord struct {
	Row int
	Col int
}

type MoveEval struct {
	Col    int8
	Board  BitmapBoard
	Winner bool
}

type GameStatus struct {
	Board  [6][7]int `json:"board"`
	Winner int       `json:"winner"`
	IsDone bool      `json:"isdone"`
	Coords []Coord   `json:"coords"`
	Info   string    `json:"info"`
	Empty  bool      `json:"empty"`
}

func (gs *GameStatus) Stringify() string {
	str, err := json.Marshal(gs)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return string(str)
}
