package typeassertions

import "fmt"

func TypeAssertions() {
	var a interface{}
	// Stores a string
	a = "Hello World"
	// Stores an integer
	a = 10
	// Type Assertion
	interfaceValue := a.(int)
	// Print the value.
	fmt.Println(interfaceValue)
}

// ❗ Will Panic => panic: interface conversion: interface {} is string, not int
func TypeAssertionsPanic() {
	var a interface{}
	// Stores a string
	a = "Hello World"
	// Stores an integer
	// a = 10
	// Type Assertion
	interfaceValue := a.(int)
	// Print the value.
	fmt.Println(interfaceValue)
}
