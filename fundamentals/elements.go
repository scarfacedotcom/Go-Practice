package main

import (
	"fmt"
	"log"
	"strings"
)

func upper(countries []string) {
	for k, _ := range countries {
		countries[k] = strings.ToUpper(countries[k])

	}

}

func main() {
	EUcountries := []string{"austria", "belgium", "bulgaria"}
	upper(EUcountries)
	fmt.Println(EUcountries)
	log.Println(EUcountries)
}
