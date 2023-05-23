package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}
type Bus struct {
	FrontSeat Passenger
}

func main() {
	peter := Passenger{"Peter", 1, false}
	fmt.Println(peter)
	var (
		scar = Passenger{Name: "scar", TicketNumber: 2, Boarded: true}
		face = Passenger{Name: "face", TicketNumber: 3}
	)
	fmt.Println(scar)
	fmt.Println(face)
}
