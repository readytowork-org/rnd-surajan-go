package type_checking

import (
	"fmt"
	"reflect"
)

func UsingReflectValueOfKind() {
	var radius1 int = 4
	var radius2 float32 = 6.512
	var radius3 float64 = 8.123456789
	var yesNo bool = true

	fmt.Println(reflect.ValueOf(radius1))
	fmt.Println(reflect.ValueOf(radius1).Kind())

	fmt.Println(reflect.ValueOf(radius2))
	fmt.Println(reflect.ValueOf(radius2).Kind())

	fmt.Println(reflect.ValueOf(radius3))
	fmt.Println(reflect.ValueOf(radius3).Kind())

	fmt.Println(reflect.ValueOf(yesNo))
	fmt.Println(reflect.ValueOf(yesNo).Kind())
}
