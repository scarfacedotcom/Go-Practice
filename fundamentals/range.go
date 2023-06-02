package main

import (
	"fmt"
)

func main() {
	fmt.Println("------ARRAYS------------")
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Println(index, value)
	}
	fmt.Println("------STRINGS------------")
	slice := []string{"scar", "face", "!"}
	for i, element := range slice {
		fmt.Println(i, element, ":")
		for _, ch := range element {
			fmt.Printf(" %q\n", ch)
		}
	}

	fmt.Println("------MAPS------------")
	person := map[string]int{"peter": 20, "scar": 30, "jay": 40}
	for name, age := range person {
		fmt.Println(name, age)
	}
}
