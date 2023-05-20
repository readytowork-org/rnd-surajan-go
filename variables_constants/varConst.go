package variables_constants

import "fmt"

func Variables() {
	// HERE, WE USED "int" not "int8" or "int64" or so on.
	// Using only "int" will have the variable size defined by the architecture the program is compiled for.
	var numOfFlavors int
	numOfFlavors = 57
	fmt.Println(numOfFlavors)

	// Declare flavorScale below:
	var flavorScale float32 = 5.8
	fmt.Println(flavorScale)

	// Constant
	const earthsGravity = 9.80665
	fmt.Println(earthsGravity)
}
