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
	return (cap(s.items) - 1) == s.top
}

func (s *Stack) IsEmpty() bool {
	return -1 == s.top
}

func (s *Stack) Push(element int) {
	s.top++
	s.items[s.top] = element
}

func (s *Stack) Pop() {
	s.items[s.top] = 0
	s.top--
}

func (s *Stack) Peek(i int) int {
	return s.items[i]
}

func (s *Stack) TopIndex() int {
	return s.top + 1
}
