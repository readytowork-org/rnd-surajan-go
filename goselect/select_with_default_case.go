package goselect

import (
	"fmt"
	"time"
)

func SelectWithDefaultCase() {
	// Create channels
	number := make(chan int)
	message := make(chan string)

	// Call goroutines that sends data to the channels.
	go channelNumb(number)
	go channelMessg(message)

	// Select and execute one channel based on availability.
	select {
	case firstChannel := <-number:
		fmt.Println("Channel Data: ", firstChannel)
	case secondChannel := <-message:
		fmt.Println("Channel Data: ", secondChannel)
	// Default Case
	default:
		fmt.Println("Wait! Channels are not ready for execution")
	}
}

// Send integer data to "number" channel
func channelNumb(number chan int) {
	// Sleep this function by 2 seconds
	time.Sleep(2 * time.Second)
	// Send data to channel only after waking up after 2 seconds.
	number <- 10
}

// Send string data to "message" channel
func channelMessg(message chan string) {
	// Sleep this function by 2 seconds
	time.Sleep(2 * time.Second)
	// Send data to channel only after waking up after 2 seconds.
	message <- "Learning Go Select"
}
