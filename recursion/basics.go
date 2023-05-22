package recursion

import "fmt"

func getFactorial(num int) int {
	if num == 0 {
		return 1
	} else {
		// Calling the function "getFactorial" inside itself.
		return num * getFactorial(num-1)
	}
}

func Recursion() {
	var num int = 3
	var factorial int = getFactorial(num)
	fmt.Println(factorial)
}
