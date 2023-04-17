package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["egss"] = 11
	shoppingList["tea"] = 3
	shoppingList["milk"] += 1
	shoppingList["egss"] += 1

	fmt.Println(shoppingList)
	delete(shoppingList, "tea")
	fmt.Println(shoppingList)

}
