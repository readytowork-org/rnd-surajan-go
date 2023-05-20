package conditionals

import "fmt"

func DeclaringVarJustBeforeUse() {
	// In If-Else
	x := 8
	y := 9
	if product := x * y; product > 60 {
		fmt.Println(product, "is greater than 60")
	}

	// In Switch
	switch season := "summer"; season {
	case "summer":
		fmt.Println("Go out and enjoy the sun!")
	}

	// NOTE: The variable "product" can only be accessed within the if-else block.
	// NOTE: The variable "season" can only be accessed within the switch block.
}
