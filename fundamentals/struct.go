// package main

// import "fmt"

// type Passenger struct {
// 	Name         string
// 	TicketNumber int
// 	Boarded      bool
// }
// type Bus struct {
// 	FrontSeat Passenger
// }

// func main() {
// 	peter := Passenger{"peter", 1, false}
// 	fmt.Println(peter)

// 	var (
// 		bill = Passenger{Name: "Bill", TicketNumber: 2}
// 		ella = Passenger{Name: "ELLA", TicketNumber: 3}
// 	)
// 	fmt.Println(bill, ella)

// 	var scar Passenger
// 	scar.Name = "scarface"
// 	scar.TicketNumber = 4
// 	fmt.Println("DECLARED", scar)

// 	peter.Boarded = true
// 	//ella.Boarded = true
// 	if peter.Boarded {
// 		fmt.Println("peter has boarded the bus")
// 	}
// 	if ella.Boarded {
// 		fmt.Println("ella is yet to board the bus")
// 	}
// 	scar.Boarded = true
// 	bus := Bus{scar}
// 	fmt.Println("front set is", bus.FrontSeat.Name)
// 	fmt.Println(bus)
// }

package main

//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates

type Coordinate struct {
	x, y int
}
type Rectangle struct {
	a Coordinate
	b Coordinate
}

func (rect Rectangle) length() int {
	return rect.a.y - rect.b.y
}
func (rect Rectangle) width() int {
	return rect.b.x - rect.a.x
}
func area(rect Rectangle) int {
	return length(rect) * width(rect)
}

//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides
func main() {

}
