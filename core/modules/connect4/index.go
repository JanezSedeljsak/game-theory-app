package connect4

func (s *State) Init(aiStart bool, isAdvanced bool) [Height][Width]int {
	s.Lock()
	defer s.Unlock()
	s.board.Init()

	s.board.Drop(2, -1)
	s.board.Drop(3, 1)
	s.board.Drop(3, -1)

	return s.board.ToMatrix()
}
