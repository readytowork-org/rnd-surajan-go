package functions

import "fmt"

// Here "int32" inside parenthesis is the type of the parameters "x" and "y".
// And "int32" outside parenthesis is the return type of the function: "multiplier".
func Multiplier(x, y int32) int32 {
	return x * y
}

func ShowResult(age int, gpa float32) string {
	ans := fmt.Sprintf("Age is %d and GPA is %.2f", age, gpa)
	return ans
}
