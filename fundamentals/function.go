package main

import "fmt"

type Coordinate struct {
	X, Y int
}

func shiftBy(x, y int, coord *Coordinate) {
	coord.X += x
	coord.Y += y
}
func main() {
	coord := Coordinate{5, 5}
	fmt.Println("Before shifting:", coord)
	shiftBy(1, 1, &coord)
	fmt.Println("After shifting:", coord)
}
