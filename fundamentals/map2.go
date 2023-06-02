package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 14
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println("First iteration", shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk has now been deleted", shoppingList)

}
