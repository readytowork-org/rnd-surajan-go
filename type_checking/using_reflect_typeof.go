package type_checking

import (
	"fmt"
	"reflect"
)

func UsingReflectTypeof() {
	var radius1 int = 4
	var radius2 float32 = 6.512
	var radius3 float64 = 8.123456789
	var yesNo bool = true

	fmt.Println(reflect.TypeOf(radius1))
	fmt.Println(reflect.TypeOf(radius2))
	fmt.Println(reflect.TypeOf(radius3))
	fmt.Println(reflect.TypeOf(yesNo))
}
