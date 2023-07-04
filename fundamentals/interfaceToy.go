package main

import "fmt"

type Toy interface {
	Play()
}

type ToyCar struct {
}

func (c ToyCar) Play() {
	fmt.Println("The car is moving")

}

type ToyAeroplane struct {
}

func (a ToyAeroplane) Play() {
	fmt.Println("The plane is moving")

}

type ToyRobot struct {
}

func (r ToyRobot) Play() {
	fmt.Println("The Robot is moving")
}
func main() {

}
