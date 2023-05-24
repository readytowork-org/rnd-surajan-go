package goroutines

import (
	"fmt"
	"time"
)

// create a function
func displayMsg(message string) {
	fmt.Println(message)
}

func MultipleGoroutines() {
	go displayMsg("Process 1")
	go displayMsg("Process 2")
	go displayMsg("Process 3")
	go displayMsg("Process 4")
	go displayMsg("Process 5")
	// Sleep the main() goroutine for 1 second.
	time.Sleep(time.Second * 1)
}
