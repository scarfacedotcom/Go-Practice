package main

import (
	"fmt"
	"log"
)

func addCountries(countriesPtr *[]string) {
	*countriesPtr = append(*countriesPtr, []string{"Croatia", "Republic of Cyprus", "Czech Republic", "Denmark", "Estonia", "Finland", "France", "Germany", "Greece", "Hungary", "Ireland", "Italy", "Latvia", "Lithuania", "Luxembourg", "Malta", "Netherlands", "Poland", "Portugal", "Romania", "Slovakia", "Slovenia", "Spain", "Sweden"}...)
}

func main() {
	EUcountries := []string{"austria", "belgium", "bulgaria"}
	addCountries(&EUcountries)
	log.Println(EUcountries)
	fmt.Println(EUcountries)
}
