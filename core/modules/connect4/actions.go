package connect4

import (
	"fmt"
	"math/rand"
)

func (s *Actions) Init(aiStart bool, isAdvanced bool) [Height][Width]int {
	s.Lock()
	defer s.Unlock()
	s.board.Init()

	if aiStart {
		s.RandomMove()
	}

	return s.board.ToMatrix()
}

func (s *Actions) PlayerDrop(col int, player int) [Height][Width]int {
	s.Lock()
	defer s.Unlock()

	st := s.board.Drop(col, player)
	fmt.Println(st, s.board, col, player)
	return s.board.ToMatrix()
}

func (s *Actions) Multiplayer() [Height][Width]int {
	return s.board.ToMatrix()
}

func (s *Actions) RandomMove() [Height][Width]int {
	moveOptions := s.board.GetOpenSpots()
	randCol := moveOptions[rand.Intn(len(moveOptions))]
	s.board.Drop(randCol, -1)
	return s.board.ToMatrix()
}