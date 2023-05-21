package arrays

import "fmt"

func ArrayWithoutElements() {
	// Create the array here
	var playerNames [5]string
	fmt.Println(playerNames)
}

func ArrayWithElements() {
	// create array below
	lottery := [5]int{5, 48, 32, 1, 6}
	fmt.Println(lottery)
}

func ArrayUsingEllipsis() {
	// This will get the compiler to determine the length of the array automatically.
	lottery := [...]int{5, 48, 32, 1, 6}
	fmt.Println(lottery)
}
