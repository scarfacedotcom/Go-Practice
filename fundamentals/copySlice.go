package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40}
	b := []int{1, 2, 3}
	copy(b, a)
	fmt.Println(b)
	copy(a, b)
	fmt.Println(a)
}
