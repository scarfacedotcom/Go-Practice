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

	var jay Passenger
	jay.Name = "Jay"
	jay.TicketNumber = 4
	fmt.Println(jay)

	scar.Boarded = false
	face.Boarded = true

	if face.Boarded {
		fmt.Println("face has boarded the bus")
	}
	if scar.Boarded {
		fmt.Println(scar.Name, "has boarded the bus")
	} else {
		fmt.Println(scar.Name, "has not boarded the bus yet")
	}

	jay.Boarded = true
	bus := Bus{jay}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
