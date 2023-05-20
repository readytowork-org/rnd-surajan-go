package fmt_package

import "fmt"

func Sprintf() {
	correctAns := "A"
	answer := fmt.Sprintf("And the correct answer is… %v!", correctAns)

	fmt.Print(answer) // Prints: And the correct answer is… A!
}

func Another_Sprintf() {
	template := "I wish I had a %v."
	pet := "puppy"

	var wish string

	// Add your code below:
	wish = fmt.Sprintf(template, pet)

	fmt.Println(wish)
	// Prints: I wish I had a puppy.
}
