package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println("your room is cleaned")
		} else {
			fmt.Println("you are dirty bro")
		}
	}

}
func main() {
	rooms := [...]Room{
		{name: "bedroom", cleaned: true},
		{name: "office", cleaned: true},
		{name: "warehouse"},
		{name: "oops"},
	}
	checkCleanliness(rooms)
	fmt.Println("performing cleaning...")
	rooms[2].cleaned = true
	rooms[3].cleaned = true
	checkCleanliness(rooms)
}
