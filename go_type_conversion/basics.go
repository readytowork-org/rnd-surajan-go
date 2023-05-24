package go_type_conversion

import "fmt"

func TypeConversion() {
	var num1 int = 4
	var num2 float32 = 4.6
	// Here, we changed the type of "num1" from "int" to "float32"
	fmt.Println(float32(num1) + num2)
	fmt.Printf("Type of num1 is %T", float32(num1))
}
