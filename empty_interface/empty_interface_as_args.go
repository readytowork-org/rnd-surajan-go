package emptyinterface

import "fmt"

// Using empty "interface{}" as parameter.
func display(i interface{}) {
	fmt.Println(i)
}

func EmptyInterfaceAsArgs() {
	a := "Hello World"
	b := 16
	c := false

	// Pass string as an argument to "display()" function.
	display(a)
	// Pass number as an argument to "display()" function.
	display(b)
	// Pass bool as an argument to "display()" function.
	display(c)
}
