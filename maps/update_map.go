package maps

import "fmt"

func UpdateMap() {
	donuts := map[string]int{
		"frosted":   10,
		"chocolate": 15,
		"jelly":     8,
	}

	// Print out all the donuts
	fmt.Println(donuts)

	// Add a new key to our "donuts" map.
	donuts["glazed"] = 12
	fmt.Println(donuts["glazed"])
	// Update the value of "jelly" key.
	donuts["jelly"] = 3
	fmt.Println(donuts)
}
