package connect4

import "reflect"

type Stack struct {
	items []interface{}
	top   int
	t     reflect.Type // current type
}

func (s *Stack) pushTypeCheck(v interface{}) {
	typ := reflect.TypeOf(v)
	if s.t != nil && typ.PkgPath()+"#"+typ.Name() != s.t.PkgPath()+"#"+s.t.Name() {
		panic("[Stack] trying to push different types to stack!")
	}
}

func (s *Stack) Init(size int) bool {
	s.items = make([]interface{}, size)
	s.top = -1
	return true
}

func (s *Stack) IsFull() bool {
	return (cap(s.items) - 1) == s.top
}

func (s *Stack) IsEmpty() bool {
	return -1 == s.top
}

func (s *Stack) Push(element interface{}) {
	if s.IsFull() {
		panic("[Stack] trying to push to full stack!")
	}

	if s.IsEmpty() {
		// set type of elements on stack based on first inserted
		s.t = reflect.TypeOf(element)
	}

	s.top++
	s.items[s.top] = element
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic("[Stack] cannot pop from empty stack!")
	}

	val := s.items[s.top]
	s.items[s.top] = 0
	s.top--
	return val
}

func (s *Stack) Peek(i int) interface{} {
	return s.items[i]
}

func (s *Stack) Count() int {
	return s.top + 1
}
