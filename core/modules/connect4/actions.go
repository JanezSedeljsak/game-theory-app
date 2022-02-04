package connect4

import "math/rand"

func (s *Actions) Init(aiStart bool, isAdvanced bool) [Height][Width]int {
	s.Lock()
	defer s.Unlock()
	s.board.Init()

	if isAdvanced {
		s.bitmap.Init()
		if aiStart {
			// best initial move is to drop in the center
			s.bitmap.MakeMove(int8(Width)/2, -1)
		}

		return s.bitmap.ToMatrix()
	}

	if aiStart {
		moveOptions := s.board.GetOpenSpots()
		randCol := moveOptions[rand.Intn(len(moveOptions))]
		s.board.Drop(randCol, -1)
	}

	return s.board.ToMatrix()
}

func (s *Actions) Mutate(board [Height][Width]int, lastCol int8) string {
	s.Lock()
	defer s.Unlock()

	s.bitmap.MakeMove(lastCol, 1)
	s.board.FromMatrix(s.bitmap.ToMatrix())
	s.board.SetLastInserted(int(s.bitmap.lastCol))

	gs := getGameStatus(s.board)
	if gs.IsDone {
		return gs.Stringify()
	}

	aiMove := CalcMove(s.bitmap)
	s.bitmap.MakeMove(aiMove.Col, -1)
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
