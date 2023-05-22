package closure

import "fmt"

func displayName() func() string {
	// Variable inside the outer function.
	name := "John Wick"

	// Anonymous function.
	return func() string {
		// Here, anonymous function can access "name" which is defined in the outer function.
		return "Hello " + name
	}
}

func Closure() {
	// Calling the outer function
	showName := displayName()

	// Calling the inner function
	fmt.Println(showName())
}
