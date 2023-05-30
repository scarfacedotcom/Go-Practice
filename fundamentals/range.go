package main

import "fmt"

func main() {
	fmt.Println("------arrays------------")
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Println(index, value)
	}
	fmt.Println("------maps------------")
	person := map[string]int{"peter": 20, "scar": 30, "jay": 40}
	for name, age := range person {
		fmt.Println(name, age)
	}
}
