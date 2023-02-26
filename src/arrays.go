package main

import "fmt"

type Product struct {
	name  string
	price float64
}

type ShoppingList = [4]Product

func printShoppingListInfo(list ShoppingList) {
	totalItems := 0
	lastItem := Product{}
	totalCost := float64(0)

	for i := 0; i < len(list); i++ {
		item := list[i]

		if item.price > 0 {
			totalItems++
			lastItem = item
			totalCost += item.price
		}
	}

	fmt.Println("==========================")
	fmt.Println("LIST INFO")
	fmt.Println("Total items:", totalItems)
	fmt.Println("Total cost:", totalCost)
	fmt.Println("Last item:", lastItem)
	fmt.Println("==========================")
}

func main() {

	var (
		p1 = Product{"iPhone", 25000}
		p2 = Product{"Mandarinka", 8.5}
		p3 = Product{"27inch screen", 6500}
		p4 = Product{"Kaiserka", 5.5}
		p5 = Product{"Mrkva", 6}
		p6 = Product{"Spaghetti", 79.9}
		p7 = Product{"Yoghurt", 19.5}
		//p8  = Product{"Teplaky", 799.99}
		//p9  = Product{"Tricko", 499.99}
		//p10 = Product{"Sveter", 1799.99}
	)

	shoppingList1 := [4]Product{p1, p2, p3}
	printShoppingListInfo(shoppingList1)

	shoppingList2 := [4]Product{p4, p5, p6}
	printShoppingListInfo(shoppingList2)

	shoppingList1[3] = p7
	printShoppingListInfo(shoppingList1)

}
