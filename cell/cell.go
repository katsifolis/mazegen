package cell

// Cell struct
type Cell struct {
	Wall    [4]bool // left-right-up-down
	X       int
	Y       int
	Visited bool
}

// Maze inits a maze duh..
func Maze(x int, y int) [][]Cell {

	maze := make([][]Cell, y)
	for i := range maze {
		maze[i] = make([]Cell, x)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			maze[i][j] = Cell{[...]bool{false, false, false, false}, i, j, false}
		}
	}

	return maze
}
