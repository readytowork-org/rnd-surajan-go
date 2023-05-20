package fmt_package

import "fmt"

func OneInput() {
	fmt.Println("What would you like for lunch?")

	// Add your code below:
	var food string
	fmt.Scan(&food)
	fmt.Printf("Sure, we can have %v for lunch.", food)
}

func TwoInputs() {
	fmt.Println("How are you doing?")

	var response1 string
	var response2 string
	// Accepts two inputs using space. Give ans like this: "not bad" with space.
	fmt.Scan(&response1)
	fmt.Scan(&response2)

	fmt.Printf("I'm %v %v", response1, response2)
}
