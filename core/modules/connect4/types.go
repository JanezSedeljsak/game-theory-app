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
	board Board
}

type Coord struct {
	Row int
	Col int
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
