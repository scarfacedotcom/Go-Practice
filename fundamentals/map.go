package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["bama"] = 10
	shoppingList["tea"] = 4
	shoppingList["milk"] += 1
	shoppingList["eggs"] += 1

	fmt.Println(shoppingList)
	delete(shoppingList, "tea")
	fmt.Println("Tea deleted, New List", shoppingList)
	fmt.Println("need", shoppingList["eggs"], "eggs")
	sugar, found := shoppingList["sugar"]
	if !found {
		fmt.Println("sugar not found in the map", sugar)
	}
	// bama, found := shoppingList["bama"]
	// if !found {
	// 	fmt.Println("no bama bro", bama)
	// } else {
	// 	fmt.Println("we currently have like", bama, "boxes of bama")
	//}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("There are", totalItems, "on the shopping list")

}
