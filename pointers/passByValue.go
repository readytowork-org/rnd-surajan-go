package pointers

import "fmt"

func addHundred(num int) {
	num += 100
}

func PassByValue() {
	x := 1
	addHundred(x)
	fmt.Println(x) // Still prints 1 instead of 101.
	// So, changing a variable's value using a function is not possible as Go is a Pass-By-Value Language.
	// This is where, Pointers come into play.
}
