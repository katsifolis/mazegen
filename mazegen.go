package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"github.com/oxilgk/mazegen/cell"
	"github.com/oxilgk/mazegen/stack"
)

func backtracer(x, y int) {

	r := rand.New(rand.NewSource(time.Now().Unix()))
	rx := r.Int() % x
	ry := r.Int() % y

	m := cell.Maze(x, y)
	s := stack.New()

	m[rx][ry].Visited = true
	initial := m[rx][ry]
	s.Push(initial)

	for s.Size() > 0 {
		curr := s.Pop().(cell.Cell)
		print(curr.X)
		print(" - ")
		println(curr.Y)
		cell.BoundCheck(x-1, y-1, &curr)
		if 

	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if m[i][j].Visited == true {
				print(".")
			}
		}
		println()
	}
}

}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	// if len(os.Args) < 3 {
	// 	println("Parameters not given")
	// 	os.Exit(0)
	// }
	// x, _ := strconv.Atoi(os.Args[1])
	// y, _ := strconv.Atoi(os.Args[2])

	x := 4
	y := 4

	backtracer(x, y)

	// pixelgl.Run(run)

}
