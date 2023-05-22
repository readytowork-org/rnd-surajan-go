package emptyinterface

import "fmt"

// Create an empty interface
// NOTE: WE USED "var" instead of "type" for creating empty interface.
var a interface{}

func EmptyInterface() {
	a = "Hello World"
	a = 10
	fmt.Println(a)
}
