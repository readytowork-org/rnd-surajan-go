package gochannels

import "fmt"

func GoChannels() {
	// Creating a channel named "numberChannel" of "int" type.
	numberChannel := make(chan int)
	fmt.Printf("Channel Type: %T\n", numberChannel)
	// The value of the channel will be a memory address.
	// It is through this address that multiple goroutines can talk and share resources to each other.
	fmt.Printf("Channel Value: %v", numberChannel)
}
