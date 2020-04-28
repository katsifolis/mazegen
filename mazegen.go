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
		for i := 0; i < 4; i++ {
			if curr.Wall[i] == false {
				println("hello there")
				s.Push(curr)
				switch i {
				case 0:
					print("in")
					if curr.Y-1 < 0 {
						break
					}
					m[curr.X][curr.Y-1].Wall[i] = true
					m[curr.X][curr.Y-1].Visited = true
					curr = m[curr.X][curr.Y-1]
					s.Push(curr)
				case 1:
					print("in")
					if curr.Y+1 > y {
						break
					}
					m[curr.X][curr.Y+1].Wall[i] = true
					m[curr.X][curr.Y+1].Visited = true
					curr = m[curr.X][curr.Y+1]
					s.Push(curr)
				case 2:
					print("in")
					if curr.X-1 < 0 {
						break
					}
					m[curr.X-1][curr.Y].Wall[i] = true
					m[curr.X-1][curr.Y].Visited = true
					curr = m[curr.X-1][curr.Y]
					s.Push(curr)
				case 3:
					print("in")
					if curr.X+1 > x {
						break
					}
					m[curr.X+1][curr.Y].Wall[i] = true
					m[curr.X+1][curr.Y].Visited = true
					curr = m[curr.X+1][curr.Y]
					s.Push(curr)
				}
			}
		}
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
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

	// x := 20
	// y := 20

	backtracer(x, y)

	// pixelgl.Run(run)

}
