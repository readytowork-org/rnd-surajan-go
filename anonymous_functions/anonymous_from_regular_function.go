package anonymousfunctions

import "fmt"

// This says that we are returning an anonymous function which is then returning an int.
func displayNumber() func() int {
	num := 1
	return func() int {
		num++
		return num
	}
}

func AnonymousFuncFromRegularFunc() {
	// Here, number will store the anonymous function and calling "number()" will return 2.
	number := displayNumber()
	fmt.Println(number())
}
