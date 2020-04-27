package main

// Cell struct
type Cell struct {
	wall    byte // -- 0x0000 first octed unused -- 0x0000 four walls
	x       int
	y       int
	visited bool
}

// Stack struct
type Stack struct {
	size  int
	top   int
	items []int
}

// InitMaze inits a maze duh..
func InitMaze(x int, y int) [][]Cell {

	maze := make([][]Cell, y)
	for i := range maze {
		maze[i] = make([]Cell, x)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			maze[i][j] = Cell{0, i, j, false}
		}
	}

	return maze
}

// StackInit inits a Stack
func StackInit(s int) Stack {

	return Stack{s, 0, make([]int, s)}
}

// Pop a value from the Stack
func (s *Stack) Pop() int {
	if s.size < 0 {
		return -1
	}
	s.top--
	r := s.items[s.top]
	return r

}

// Push pushes a value to the Stack
func (s *Stack) Push(v int) bool {

	if s.top == s.size {
		return false

	}
	s.items[s.top] = v
	s.top++
	return true

}
