package gogenerics

import "fmt"

func GenericsAnyType() {
	var num int = 4
	var name string = "Aaron Swartz"

	message(num)
	message(name)
}

// Using [T any], this function will accept "any" type.
func message[T any](msg T) {
	fmt.Println(msg)
}
