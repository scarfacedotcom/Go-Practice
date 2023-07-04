package main

import "fmt"

type Human struct {
	Firstname string
	Lastname string
	Age int
	Country string
}

type DomesticAnimal interface {
	ReceiveAffection(from Human)
	GiveAffecttion(to Human)
}

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

func (c Cat) ReceiveAffection(from Human) {
fmt.Printf("The cat name %s has recieved affection from Human name %s\n", c.Name, from.Firstname)	
}

func (c Cat) GiveAffection(from Human) {
	fmt.Printf("The cat name %s has recieved affection from Human name %s\n", c.Name, from.Firstname)	
}
func ()  {
	
}