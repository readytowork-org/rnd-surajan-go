package pointers

import "fmt"

func addHundredByRef(numPtr *int) {
	*numPtr += 100
}

func PassByReference() {
	x := 1
	addHundredByRef(&x)
	fmt.Println(x) // Prints 101
}

func brainwash(saying *string) {
	*saying = "Beep Boop."
}

func AnotherPassByReference() {
	greeting := "Hello there!"

	brainwash(&greeting)

	fmt.Println("greeting is now:", greeting)
}
