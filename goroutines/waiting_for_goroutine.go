package goroutines

import (
	"fmt"
	"time"
)

// create a function
func show(message string) {
	fmt.Println(message)
	// to sleep the process for 1 second
	time.Sleep(time.Second * 1)
}

func WaitingForGoroutine() {
	// call function with goroutine
	go show("Process 1")
	show("Process 2")
}
