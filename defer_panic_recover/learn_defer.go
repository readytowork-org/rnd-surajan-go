package defer_panic_recover

import "fmt"

func LearnDefer() {
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("This will be displayed at first")
}
