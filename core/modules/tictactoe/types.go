package tictactoe

import (
	"encoding/json"
	"fmt"
	"sync"
)

const Size = 3

type dp struct {
	Memo map[int]Response
}

type Coord struct {
	Row int
	Col int
}

type Response struct {
	Coords Coord
	Value  int
}

type Actions struct {
	sync.Mutex
	board Board
}

type GameStatus struct {
	Board  [3][3]int `json:"board"`
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
