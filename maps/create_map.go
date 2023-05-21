package maps

import "fmt"

func CreateMap() {
	// Creating a map using "make".
	orders := make(map[string]float32)
	fmt.Println(orders)
	// Creating a map with values.
	donuts := map[string]int{
		"frosted":   10,
		"chocolate": 15,
		"jelly":     8,
	}
	fmt.Println(donuts)
}
