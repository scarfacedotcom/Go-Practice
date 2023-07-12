package main

import "fmt"

func main() {
	hotelName := "Go Dev Hotel"
	s := hotelName[0,6]
	fmt.Println(s)
	hotelName = "Java Hotel"
	fmt.Println(s)
}
