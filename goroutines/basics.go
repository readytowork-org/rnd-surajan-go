package goroutines

import "fmt"

func display(msg string) {
	fmt.Println(msg)
}

func BasicGoroutine() {
	// Call the goroutine
	go display("First")
	// Normal function call
	display("Second")
}
