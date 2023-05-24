package gogenerics

import (
	"fmt"
	"reflect"
)

func GenericsParameterizedTypes() {
	var radius1 int = 4
	var radius2 float32 = 6.512
	var radius3 float64 = 8.123456789

	circumf(radius1)
	circumf(radius2)
	circumf(radius3)
}

type Radius interface {
	int | float32 | float64
}

func circumf[R Radius](radius R) {
	c := 2 * 3 * radius
	// "reflect.TypeOf(value)" will return the data type that the "value" holds.
	fmt.Println(reflect.TypeOf(radius))
	fmt.Println("Circumference is:", c)
}
