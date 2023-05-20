package variables_constants

import "fmt"

func DynamicTyping() {
	daysOnVacation := 6

	var hoursInDay = 24

	fmt.Println("You have spent", daysOnVacation*hoursInDay, "hours on vacation.")
	// Prints: "You have spent 144 hours on vacation."
}
