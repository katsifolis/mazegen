package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a uint16 = 0x0001
	fmt.Printf("%x\n", a&0x1000)

	if len(os.Args) < 3 {
		println("Parameters not given")
		os.Exit(0)
	}
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])

	println(x, y)
	m := cell.InitMaze(x, y)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Printf("%c ", m[i][j])
		}
		println("")
	}
}
