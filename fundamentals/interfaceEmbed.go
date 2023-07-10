package main

import "fmt"

type Animal interface {
	Eat()
	Sleep()
}
type Dog struct {
}

func (d Dog) Eat() {
	fmt.Println("THe Dog is eating")
}

func (d Dog) Sleep() {
	fmt.Println("THe Dog is sleeping")
}

type Bird struct {
}

func (b Bird) Eat() {
	fmt.Println("The bird is eating")
}

func (b Bird) Sleep() {
	fmt.Println("The bird is sleeping")
}

type Flying interface {
	Fly()
}

type FlyingAnimal interface {
	Animal
	Flying
}

type Eagle struct {
}

func (e Eagle) Eat() {
	fmt.Println("The eagle is eating")
}

func (e Eagle) Sleep() {
	fmt.Println("the eagle is sleeping")
}

func (e Eagle) Fly() {
	fmt.Println("the eagle is flying")
}

func main() {
	dog := Dog{}
	bird := Bird{}
	eagle := Eagles{}

	dog.Eat()
	bird.Sleep()
	eagle.Eat()
	eagle.Fly()
}
