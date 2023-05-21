package loops

import "fmt"

func ForLoopAsWhileLoop() {
	number := 0 // Initialize a variable to be used inside the loop
	for number < 5 {
		fmt.Println(number)
		number++ // Update the variable being used
	}
}
