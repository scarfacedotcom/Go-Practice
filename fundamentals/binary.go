package main

import (
	"fmt"
	"math/rand"
)

func main() {
	hotelName := "scar hotel"

	fmt.Println("Hotel:" + hotelName)
	var roomsAvailable int
	var rooms, roomsOccupaied int = 100, rand.Intn(100)
	roomsAvailable = rooms - roomsOccupaied
	fmt.Println(roomsAvailable, "rooms Available")

}
