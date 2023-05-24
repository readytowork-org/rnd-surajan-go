package gogenerics

import "fmt"

func GenericsEitherOr() {
	var radius1 int = 4
	var radius2 float32 = 6.512

	circumference(radius1)
	circumference(radius2)
}

// Here, [r int | float32] is a Generic telling that "r" can be either "int" or "float".
func circumference[r int | float32](radius r) {
	// NOTE: Due to varying types, we could not use "3.14" as the value of pi.
	c := 2 * 3 * radius
	fmt.Println("Circumference is:", c)
}
