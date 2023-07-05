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
func ()  {
	
}