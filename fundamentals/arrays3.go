package main

import "fmt"

func main() {
	myArray := [2]int{4, 6}
	for index, element := range myArray {
		fmt.Printf("element at index %d is %d\n", index, element)
	}
	myArray2 := [2]int{6, 7}
	for i := 0; i < len(myArray2); i++ {
		fmt.Printf("element at index %d is %d\n", i, myArray2[i])
	}
}
