package main

import "log"

type Person struct {
	name   string
	age    int
	height int
}

func (person *Person) changeName(newName string) {
	person.name = newName
}
func (person Person) changeAge(newAge int) {
	person.age = newAge
}

func main() {
	person := Person{name: "peter", age: 90, height: 44}
	person.changeName("scarface")
	log.Println("The new name now is", person.name)
	person.changeAge(40)
	log.Println("The new age now is", person.age)
}
