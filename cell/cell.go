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

// BoundCheck checks for bounds in the maze
func BoundCheck(maxx, maxy int, c *Cell) {

	if c.X == 0 {
		c.Wall[2] = true
		if c.Y == 0 {
			c.Wall[0] = true
		} else if c.Y == maxy {
			c.Wall[1] = true
		}
	} else if c.Y == 0 {
		c.Wall[0] = true
		if c.X == maxx {
			c.Wall[3] = true
		}
	} else if c.X == maxx {
		c.Wall[3] = true
		if c.Y == maxy {
			c.Wall[1] = true
		}
	} else if c.Y == maxy {
		c.Wall[1] = true
	}

	// if c.X == 0 && c.Y == 0 {
	// 	c.Wall[0], c.Wall[2] = true, true
	// } else if c.X == 0 && c.Y == maxy {
	// 	c.Wall[1], c.Wall[2] = true, true
	// } else if c.X == maxx && c.Y == 0 {
	// 	c.Wall[3], c.Wall[0] = true, true
	// } else if c.X == maxx && c.Y == maxy {
	// 	c.Wall[1], c.Wall[3] = true, true
	// } else if c.X == 0 {
	// 	c.Wall[2] = true
	// } else if c.X == maxx {
	// 	c.Wall[1] = true
	// } else if c.Y == 0 {
	// 	c.Wall[0] = true
	// } else if c.Y == maxy {
	// 	c.Wall[3] = true
	// }
}
