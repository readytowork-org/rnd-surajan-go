package structs

import "fmt"

type Triangle struct {
	height float32
	base   float32
}

// Checkpoint 1 code goes here
func (triangle Triangle) area() float32 {
	return (triangle.base * triangle.height) / 2
}

func FunctionsOfStruct() {
	triangle := Triangle{10, 4}
	fmt.Println(triangle)

	// Call the function here
	triangleArea := triangle.area()
	fmt.Println(triangleArea)
}
