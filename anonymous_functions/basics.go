package anonymousfunctions

import "fmt"

func AnonymousFunc() {
	var greet = func() {
		fmt.Println("Hello World")
	}
	greet()
}

func AnonymousFuncWithArgsAndReturn() {
	var sum = func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(1, 2))
}

func AnonymousFuncUsingShortOperator() {
	sum := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(1, 2))
}
