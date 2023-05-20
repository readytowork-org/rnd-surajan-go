package variables_constants

import "fmt"

func DefaultValues() {
	var emptyInt int8

	var emptyFloat float32

	var emptyString string

	// We can print more than one variable or literals by appending commas ","
	fmt.Println(emptyInt, emptyFloat, emptyString)
	// Prints: "0 0" as emptyString is ""
}
