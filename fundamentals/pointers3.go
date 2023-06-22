package main

import "fmt"

type Hotel struct {
	Name string
	*Country
}

type Country struct {
	Name        string
	CapitalCity string
}

func main() {
	hotel := Hotel{
		Name:    "Scar Face Hotel",
		Country: &Country{Name: "Canada"},
	}
	fmt.Println(hotel.Country.Name)

}
