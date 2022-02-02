package connect4

import (
	"math/rand"
)

func (s *Actions) Init(aiStart bool, isAdvanced bool) [Height][Width]int {
	s.Lock()
	defer s.Unlock()

	s.board.Init()
	if aiStart {
		if isAdvanced {
			// best initial move is to drop in the center
			s.board.Drop(Width/2, -1)
		} else {
			moveOptions := s.board.GetOpenSpots()
			randCol := moveOptions[rand.Intn(len(moveOptions))]
			s.board.Drop(randCol, -1)
		}
	}

	return s.board.ToMatrix()
}

func (s *Actions) Mutate(board [Height][Width]int, lastCol int) string {
	s.Lock()
	defer s.Unlock()

	s.board.FromMatrix(board)
	s.board.SetLastInserted(lastCol)
	gs := getGameStatus(s.board)
	if gs.IsDone {
		return gs.Stringify()
	}

	aiMove := CalcMove(s.board)
	s.board.Drop(int(aiMove.Col), -1)

	gs = getGameStatus(s.board)
	return gs.Stringify()
}

func (s *Actions) Multiplayer(board [Height][Width]int, lastCol int) string {
	s.Lock()
	defer s.Unlock()

	s.board.FromMatrix(board)
	s.board.SetLastInserted(lastCol)

	gs := getGameStatus(s.board)
	return gs.Stringify()
}

func (s *Actions) RandomMove(board [Height][Width]int, lastCol int) string {
	s.Lock()
	defer s.Unlock()

	gs := getGameStatus(s.board)
	if gs.IsDone {
		return gs.Stringify()
	}

	s.board.FromMatrix(board)
	s.board.SetLastInserted(lastCol)
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

func getGameStatus(board Board) GameStatus {
	var isDraw bool = board.IsDone()
	var gs GameStatus = board.CheckWinner()
	var isDone bool = isDraw || gs.Winner != 0

	return GameStatus{board.ToMatrix(), gs.Winner, isDone, gs.Coords}
}
