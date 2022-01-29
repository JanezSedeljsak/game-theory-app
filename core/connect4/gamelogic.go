package connect4

func (s *State) Init(aiStart bool, isAdvanced bool) [][]bool {
	s.Lock()
	defer s.Unlock()
	board := make([][]bool, 6)
	for i := 0; i < Height; i++ {
		board[i] = make([]bool, 7)
	}

	for i := 0; i < Width; i++ {
		board[Height-1][i] = true
	}

	return s.board
}
