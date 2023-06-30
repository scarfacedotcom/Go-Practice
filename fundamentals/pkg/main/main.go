package main

import (
	"coursecontent/fundamentals/pkg/display"
	"coursecontent/fundamentals/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("hello from display")
	msg.Exciting("an exciting message")
}
