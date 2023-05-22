package typeassertions

import "fmt"

func AvoidPanicInTypeAssertions() {
	var a interface{}
	// Stores a string
	a = "Hello World"
	// Type Assertion returns two values: The evaluated value and a boolean value that checks whether the type matches or not.
	interfaceValue, status := a.(int)
	fmt.Println(interfaceValue)
	fmt.Println(status)
}
