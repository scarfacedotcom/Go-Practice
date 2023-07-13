package main

import "fmt"

func addSolidity(languages []string) {
	languages = append(languages, "GO")
	fmt.Println("in function capacity", cap(languages))
}
func main() {
	languages := []string{"go", "java", "PHP"}
	fmt.Println("Capacity :", cap(languages))
}
