package connect4

import (
	"encoding/json"
	"fmt"
	"sync"
)

const Height int = 6
const Width int = 7

type Actions struct {
	sync.Mutex
	board   Board       // used for random AI and multiplayer
	bitmap  BitmapBoard // used for negamax (main algorithm)
	aiStart bool
}

type dp struct {
	Memo     map[uint64]int8
	MaxDepth int8
}

type Coord struct {
	Row int
	Col int
}

// Seperate from Coord to preserve memory within tree evaluation
type MiniMaxState struct {
	Col   int8
	Value int8
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
}

func (gs *GameStatus) Stringify() string {
	str, err := json.Marshal(gs)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(str)
}
