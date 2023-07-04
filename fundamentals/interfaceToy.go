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

// func PlayWithToy(t Toy) {
// 	t.Play()
// }
func main() {
	car := ToyCar{}
	airplane := ToyAeroplane{}
	robot := ToyRobot{}

	toys := []Toy{car, airplane, robot}

	for _, toy := range toys {
		toy.Play()
	}

	// PlayWithToy(car)
	// PlayWithToy(airplane)
	// PlayWithToy(robot)
	//fmt.Println(car)
}
