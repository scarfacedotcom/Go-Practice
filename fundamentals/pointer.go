package main

import "fmt"

func main() {
	x := 42
	y := &x
	fmt.Println(*y)
}
