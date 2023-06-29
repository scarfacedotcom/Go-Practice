package main

import "fmt"

type Coordinate struct {
	X, Y int
}

func (coord *Coordinate) shiftBy(x, y int) {
	coord.X += x
	coord.Y += y

}
func main() {
	coord := Coordinate{5, 5}
	fmt.Println("Before shifting:", coord)
	coord.shiftBy(1, 1)
	fmt.Println("After shifting:", coord)

}
