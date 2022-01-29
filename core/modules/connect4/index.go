package connect4

import "fmt"

func (s *State) Init(aiStart bool, isAdvanced bool) [][]bool {
	s.Lock()
	defer s.Unlock()
	s.board = make([][]bool, 6)
	for i := 0; i < Height; i++ {
		s.board[i] = make([]bool, 7)
	}

	for i := 0; i < Width; i++ {
		s.board[Height-1][i] = true
	}

	fmt.Println(s.board)
	return s.board
}
