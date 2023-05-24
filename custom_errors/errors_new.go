package custom_errors

// Importing "errors" package.
import (
	"errors"
	"fmt"
)

func checkName(name string) error {
	// Creating a new Error.
	newError := errors.New("Invalid Programming Language")

	if name != "Go" {
		return newError
	} else {
		// "nil" is used in Go instead of "null" or "undefined" like in JS.
		return nil
	}
}

func ErrorsNew() {
	name := "JavaScript"
	err := checkName(name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Valid Programming Language")
	}
	// Throwing an Error will not terminate the program immediately.
	// The code below also runs.
	fmt.Println("Continue Playing with this code")
}
