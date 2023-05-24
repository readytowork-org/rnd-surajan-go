package custom_errors

import "fmt"

func divide(num1, num2 int) error {
	// Create a formatted Error.
	divideError := fmt.Errorf("Division Error: %d / %d\nCannot divide by zero.", num1, num2)
	if num2 == 0 {
		return divideError
	} else {
		return nil
	}
}

func Errorf() {
	num1, num2 := 3, 0
	newErr := divide(num1, num2)
	if newErr != nil {
		fmt.Println(newErr)
	} else {
		fmt.Println("Division Successful")
	}
}
