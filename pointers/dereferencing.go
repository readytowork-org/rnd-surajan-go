package pointers

import "fmt"

func PointerDereferencing() {
	star := "Polaris"

	starAddress := &star

	// Add your code below:
	*starAddress = "Sirius"

	// It should print "Sirius" rather than "Polaris".
	fmt.Println("The actual value of star is", star)
}
