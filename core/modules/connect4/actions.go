package connect4

import (
	"fmt"
	"math/rand"
	"time"
)

func (s *Actions) Init(aiStart bool, isAdvanced bool) [Height][Width]int {
	s.Lock()
	defer s.Unlock()
	s.board.Init()

	if isAdvanced {
		s.bitmap.Init()
		if aiStart {
			// best initial move is to drop in the center
			s.board.Drop(Width/2, -1)
			s.bitmap.MakeMove(int8(Width) / 2)
		}
	} else {
		moveOptions := s.board.GetOpenSpots()
		randCol := moveOptions[rand.Intn(len(moveOptions))]
		s.board.Drop(randCol, -1)
	}

	return s.board.ToMatrix()
}

func (s *Actions) Mutate(board [Height][Width]int, column int8) string {
	s.Lock()
	defer s.Unlock()

	s.bitmap.MakeMove(column)
	s.board.Drop(int(column), 1)

	gs := getGameStatus(s.board)
	if gs.IsDone {
		return gs.Stringify()
	}

	// calculate depth search
	movesMade := s.board.CountMoves()
	var depth int8 = 18
	if 6 <= movesMade && movesMade < 8 {
		depth = 20
	} else if 8 <= movesMade && movesMade < 14 {
		depth = 24
	} else if movesMade >= 14 {
		depth = 42 - movesMade
	}

	start := time.Now()
	column, score := CalcMove(s.bitmap, depth)
	elapsed := time.Since(start)
	fmt.Printf("Moves made: %d, Elapsed: %s, Estimation: %d\n", movesMade, elapsed, score)

	s.bitmap.MakeMove(column)
	s.board.Drop(int(column), -1)

	gs = getGameStatus(s.board)
	return gs.Stringify()
}

func (s *Actions) Multiplayer(board [Height][Width]int, column int) string {
	s.Lock()
	defer s.Unlock()

	s.board.FromMatrix(board)
	s.board.SetLastInserted(column)

	gs := getGameStatus(s.board)
	return gs.Stringify()
}

func (s *Actions) RandomMove(board [Height][Width]int, column int) string {
	s.Lock()
	defer s.Unlock()

	gs := getGameStatus(s.board)
	if gs.IsDone {
		return gs.Stringify()
	}

	s.board.FromMatrix(board)
	s.board.SetLastInserted(column)
	gs = getGameStatus(s.board)
	if gs.IsDone {
		return gs.Stringify()
	}

	moveOptions := s.board.GetOpenSpots()
	randCol := moveOptions[rand.Intn(len(moveOptions))]
	s.board.Drop(randCol, -1)

	gs = getGameStatus(s.board)
	return gs.Stringify()
}
