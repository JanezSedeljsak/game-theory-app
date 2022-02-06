package connect4

import (
	"math/rand"
	"time"
)

func (s *Actions) Init(aiStart bool, isAdvanced bool) [Height][Width]int {
	s.Lock()
	defer s.Unlock()
	s.board.Init()
	s.History.Init(21)

	if aiStart {
		col := Ternary(isAdvanced, Width/2, rand.Intn(Width)).(int)
		s.board.Drop(col, -1)
	}

	return s.board.ToMatrix()
}

func (s *Actions) Mutate(board [Height][Width]int, column int) string {
	s.Lock()
	defer s.Unlock()

	s.logHistoryAndUpdate(board, column)
	gs := getGameStatus(s.board, "")
	if gs.IsDone {
		return gs.Stringify()
	}

	// calculate depth search
	movesMade := s.board.CountMoves()
	var depth int8 = 18
	if 8 <= movesMade && movesMade < 16 {
		depth = 23
	} else if movesMade >= 16 {
		depth = 42 - movesMade
	}

	start := time.Now()
	aiCol, score := CalcMove(s.board, depth)
	elapsed := time.Since(start)
	s.board.Drop(int(aiCol), -1)

	toastMessage := GetInfoMessage(score, elapsed, movesMade)
	gs = getGameStatus(s.board, toastMessage)
	return gs.Stringify()
}

func (s *Actions) Multiplayer(board [Height][Width]int, column int) string {
	s.Lock()
	defer s.Unlock()

	s.logHistoryAndUpdate(board, column)
	gs := getGameStatus(s.board, "")
	return gs.Stringify()
}

func (s *Actions) RandomMove(board [Height][Width]int, column int) string {
	s.Lock()
	defer s.Unlock()

	s.logHistoryAndUpdate(board, column)
	gs := getGameStatus(s.board, "")
	if gs.IsDone {
		return gs.Stringify()
	}

	moveOptions := s.board.GetOpenSpots()
	randCol := moveOptions[rand.Intn(len(moveOptions))]
	s.board.Drop(randCol, -1)

	gs = getGameStatus(s.board, "")
	return gs.Stringify()
}

func (s *Actions) PrevMove() string {
	if s.History.IsEmpty() {
		gs := GameStatus{Empty: true}
		return gs.Stringify()
	}

	prev := s.History.Pop().(Board)
	s.board = prev

	gs := getGameStatus(s.board, "")
	return gs.Stringify()
}

func (s *Actions) logHistoryAndUpdate(board [Height][Width]int, column int) {
	s.History.Push(s.board)
	s.board.FromMatrix(board)
	s.board.SetLastInserted(column)
}
