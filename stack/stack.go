package stack

// Stack struct
type Stack struct {
	top  *node
	size int
}

type node struct {
	value interface{}
	prev  *node
}

// New func
func New() *Stack {
	return &Stack{nil, 0}
}

// Size of stack
func (s *Stack) Size() int {
	return s.size
}

// Push pushes a value to the Stack
func (s *Stack) Push(v interface{}) {
	n := &node{v, s.top}
	s.top = n
	s.size++

}

// Pop a value from the Stack
func (s *Stack) Pop() interface{} {

	if s.size < 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.size--
	return n.value
}
