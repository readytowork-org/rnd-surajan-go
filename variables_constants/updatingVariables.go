package variables_constants

import "fmt"

func NormalUpdating() {
	var basketTotal float64
	carrotPrice := 0.75

	// Here, basketTotal has 0.0 value by default, so adding a new float is even easier.
	basketTotal = basketTotal + carrotPrice
	fmt.Println(basketTotal) // Prints: 0.75
}

func UpdatingUsingShorthand() {
	var basketTotal float64
	carrotPrice := 0.75

	basketTotal += carrotPrice
	fmt.Println(basketTotal) // Prints: 0.75
}
