package goselect

import "fmt"

func GoSelect() {
	// Create channels
	number := make(chan int)
	message := make(chan string)
	yesNo := make(chan bool)

	// Call goroutines that sends data to the channels.
	go channelNumber(number)
	go channelMessage(message)
	go channelYesNo(yesNo)

	// Select and execute one channel based on availability.
	select {
	case firstChannel := <-number:
		fmt.Println("Channel Data: ", firstChannel)
	case secondChannel := <-message:
		fmt.Println("Channel Data: ", secondChannel)
	case thirdChannel := <-yesNo:
		fmt.Println("Channel Data: ", thirdChannel)
	}
}

// Send integer data to "number" channel
func channelNumber(number chan int) {
	number <- 10
}

// Send string data to "message" channel
func channelMessage(message chan string) {
	message <- "Learning Go Select"
}

// Send bool data to "yesNo" channel
func channelYesNo(yesNo chan bool) {
	yesNo <- false
}
