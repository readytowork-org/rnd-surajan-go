package fmt_package

import "fmt"

func Sprintln_VS_Sprint() {
	// Sprint
	grade := "100"
	compliment := "Great job!"
	teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)

	fmt.Print(teacherSays)
	// Prints: You scored a 100 on the test! Great job!

	// Sprintln
	quote := fmt.Sprintln("Look ma,", "no spaces!")
	fmt.Print(quote) // Prints: Look ma, no spaces!
}
