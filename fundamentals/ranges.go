package main

import (
	"fmt"
)

func main() {
	slice := []string{"Hello", "world", "!"}
	for i, element := range slice {
		fmt.Println(i, element, ":")
		for _, ch := range element {
			fmt.Printf("  %q\n", ch)
		}
	}
	number := []int{1, 2, 3, 4, 5}
	for i, element := range number {
		fmt.Printf("Index: %d, Value: %d\n", i, element)
	}
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println("index", i, "value", v)
	}
	slices := []string{"petr", "scar", "face"}
	for i, v := range slices {
		fmt.Println("index", i, "string", v)
		for _, ch := range v {
			fmt.Printf("  %q\n", ch)
		}
	}
}
