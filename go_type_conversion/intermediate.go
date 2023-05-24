package go_type_conversion

import (
	"fmt"
	"math"
)

func TypeConversionIntermediate() {
	var x, y int = 4, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println(f)
}
