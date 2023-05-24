package gochannels

import "fmt"

func SendReceiveData() {
	// Create channels
	number := make(chan int)
	message := make(chan string)

	// Call a "goroutine" to send data to the channels
	go channelData(number, message)

	// Receive data from the channels
	fmt.Println("Number Channel Data:", <-number)
	fmt.Println("Message Channel Data:", <-message)
}

func channelData(number chan int, message chan string) {
	// Send data to the channels
	number <- 10
	message <- "Hello Go Channel"
}
