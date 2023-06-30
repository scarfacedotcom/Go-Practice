package main

import (
	"fmt"
	"log"
)

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
func main() {
	result := sum(1, 2, 3, 4, 5, 6)
	log.Println(result)
	//.Printf(result)
	fmt.Println(result)
}
