package emptyinterface

import "fmt"

// Using "i... interface{}" as parameter.
func display1(i ...interface{}) {
	fmt.Println(i)
}

func EmptyInterfaceAsAnyNumOfArgs() {
	a := "Hello World"
	b := 16
	c := false

	// Pass only one argument to "display1()" function.
	display1(a)
	// Pass two arguments of different data types to "display1()" function.
	display1(a, b)
	// Pass three arguments of different data types to "display1()" function.
	display1(a, b, c)
}
