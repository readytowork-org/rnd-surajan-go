package pointers

import "fmt"

func PointerBasics() {
	star := "Polaris"

	// Here, &star is the address of "star" variable.
	/*
		And, "starAddress" is a pointer which can store address of a variable whose
		data type is "string".
	*/
	var starAddress *string = &star
	fmt.Println("The address of star is ", starAddress)
}
