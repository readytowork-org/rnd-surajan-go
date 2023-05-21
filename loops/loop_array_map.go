package loops

import "fmt"

func LoopArrayMap() {
	// ARRAY
	menu := []string{"Hamburgers", "Cheeseburgers", "Fries"}

	fmt.Println("The menu:")
	// WRITE LOOP GOING THROUGH MENU BELOW THIS LINE
	for index, value := range menu {
		fmt.Printf("Item Name: %v and Item Number: %d\n", value, index)
	}

	// MAP
	orders := map[string]string{
		"John":   "Cheeseburgers",
		"Janet":  "Hamburgers",
		"Jordan": "Fries",
	}
	// A late order
	orders["James"] = "Chicken Sandwich"

	fmt.Println("\nThe friend's orders:")
	// WRITE LOOP GOING THROUGH ORDERS BELOW THIS LINE
	for key, value := range orders {
		fmt.Printf("Order for: %v and Order Name: %v\n", key, value)
	}
}
