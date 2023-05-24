package custom_errors

import "fmt"

type DivisionByZero struct {
	message string
}

// define Error() method on the struct
func (z DivisionByZero) Error() string {
	return "Number Cannot Be Divided by Zero"
}

func division(n1 int, n2 int) (int, error) {

	if n2 == 0 {
		// Here, "&DivisionByZero{}" is an instance of "DivisionByZero" struct.
		// NOTE: We are not calling Error() anywhere directly. We can access Error()'s return value using the "&DivisionByZero{}" instance.
		return 0, &DivisionByZero{}
	} else {
		return n1 / n2, nil
	}
}

func ErrorUsingInterface() {

	number1 := 15
	// change the value of number2 to get different result
	number2 := 0

	result, err := division(number1, number2)

	// check if error occur or not
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %d", result)
	}
}
