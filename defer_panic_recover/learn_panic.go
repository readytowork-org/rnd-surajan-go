package defer_panic_recover

import "fmt"

func divide(num1, num2 int) {
	if num2 == 0 {
		panic("Cannot divide by zero.")
	} else {
		fmt.Printf("Result: %v\n", num1/num2)
	}
}

func LearnPanic() {
	divide(4, 2)
	divide(4, 0)
	// This will not run as the program terminates in the 2nd "divide()" function.
	divide(5, 2)
}
