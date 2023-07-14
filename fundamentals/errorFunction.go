package main

import "fmt"

func addSolidity(languages *[]string) {
	*languages = append(*languages, "Solidity")
	fmt.Println("in function capacity", cap(*languages))
}
func main() {
	languages := []string{"go", "java", "PHP"}
	fmt.Println("Capacity :", cap(languages))

	addSolidity(&languages)
	fmt.Println("Capacity :", cap(languages))
	fmt.Println(languages)
}
