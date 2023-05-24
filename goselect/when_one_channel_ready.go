package goselect

import (
	"fmt"
	"time"
)

func WhenOneChannelReady() {
	// Create channels
	number := make(chan int)
	message := make(chan string)

	// Call goroutines that sends data to the channels.
	go sendChannelNumber(number)
	go sendChannelMessage(message)

	// Select and execute one channel based on availability.
	select {
	case firstChannel := <-number:
		fmt.Println("Channel Data: ", firstChannel)
	case secondChannel := <-message:
		fmt.Println("Channel Data: ", secondChannel)
	}
}

// Send integer data to "number" channel
func sendChannelNumber(number chan int) {
	number <- 10
}

// Send string data to "message" channel
func sendChannelMessage(message chan string) {
	// Sleep this function by 2 seconds
	time.Sleep(2 * time.Second)
	// Send data to channel only after waking up after 2 seconds.
	message <- "Learning Go Select"
}
