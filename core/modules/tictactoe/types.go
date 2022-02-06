package tictactoe

import (
	"encoding/json"
	"log"
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

// Seperate from Coord to preserve memory within tree evaluation
type Response struct {
	Row   int8
	Col   int8
	Value int8
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
		log.Fatal(err)
		return ""
	}

	return string(str)
}
