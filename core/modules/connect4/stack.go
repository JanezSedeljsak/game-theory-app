package connect4

type Stack struct {
	items []int
	top   int
}

func (s *Stack) Init(size int) bool {
	s.items = make([]int, size)
	s.top = -1
	return true
}

func (s *Stack) IsFull() bool {
	if (cap(s.items) - 1) == s.top {
		return true
	}
	return false
}

func (s *Stack) IsEmpty() bool {
	if -1 == s.top {
		return true
	}
	return false
}

func (s *Stack) Push(element int) {
	s.top++
	if s.top == -1 {
		s.items[0] = element
	} else {
		s.items[s.top] = element
	}
}

func (s *Stack) Pop() {
	s.items[s.top] = 0
	s.top--
}

func (s *Stack) Peek(i int) int {
	return s.items[i]
}
