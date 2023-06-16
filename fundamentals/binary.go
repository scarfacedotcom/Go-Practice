package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	hotelName := "scar hotel"

	fmt.Println("Hotel:" + hotelName)
	var roomsAvailable int
	rand.Seed(time.Now().UTC().UnixNano())
	var rooms, roomsOccupaied int = 100, rand.Intn(100)
	roomsAvailable = rooms - roomsOccupaied
	fmt.Println(roomsAvailable, "rooms Available")

}
