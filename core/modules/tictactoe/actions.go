package tictactoe

import "math/rand"

func (s *Actions) Init(aiStart bool, isAdvanced bool) [3][3]int {
	s.Lock()
	defer s.Unlock()

	s.board.Init()
	if aiStart {
		if isAdvanced {
			aiMove := CalcMove(s.board)
			s.board.Set(int(aiMove.Row), int(aiMove.Col), -1)
		} else {
			moveOptions := s.board.GetOpenSpots()
			randCoord := moveOptions[rand.Intn(len(moveOptions))]
			s.board.Set(randCoord.Row, randCoord.Col, -1)
		}
	}

	return s.board.ToMatrix()
}

func (s *Actions) Mutate(board [Size][Size]int) string {
	s.Lock()
	defer s.Unlock()

	gs := getGameStatus(s.board)
	if gs.IsDone {
		return gs.Stringify()
	}

	s.board.FromMatrix(board)
	gs = getGameStatus(s.board)
	if gs.IsDone {
		return gs.Stringify()
	}

	aiMove := CalcMove(s.board)
	s.board.Set(int(aiMove.Row), int(aiMove.Col), -1)
	gs = getGameStatus(s.board)
	return gs.Stringify()
}

func (s *Actions) Multiplayer(board [Size][Size]int) string {
	s.Lock()
	defer s.Unlock()

	s.board.FromMatrix(board)
	gs := getGameStatus(s.board)
	return gs.Stringify()
}

func (s *Actions) RandomMove(board [Size][Size]int) string {
	s.Lock()
	defer s.Unlock()

	gs := getGameStatus(s.board)
	if gs.IsDone {
		return gs.Stringify()
	}

	s.board.FromMatrix(board)
	gs = getGameStatus(s.board)
	if gs.IsDone {
		return gs.Stringify()
	}

	moveOptions := s.board.GetOpenSpots()
	randCoord := moveOptions[rand.Intn(len(moveOptions))]
	s.board.Set(randCoord.Row, randCoord.Col, -1)
	gs = getGameStatus(s.board)
	return gs.Stringify()
}

func getGameStatus(board Board) GameStatus {
	var isDraw bool = board.IsDone()
	var gs GameStatus = board.CheckWinner()
	var isDone bool = isDraw || gs.Winner != 0

	return GameStatus{board.ToMatrix(), gs.Winner, isDone, gs.Coords}
}
