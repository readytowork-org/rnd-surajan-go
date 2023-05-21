package structs

import "fmt"

type Triangle1 struct {
	height float32
	base   float32
}

// Struct Function that utilizes a "triangle" pointer that points to "Triangle1" struct.
// So, now we can change the value of "base" permanently.
func (triangle *Triangle1) updateBase() {
	triangle.base = 13
}

func PointersToStruct() {
	triangle := Triangle1{10, 4}
	fmt.Println(triangle)

	// Using Pointers with Struct Functions
	triangle.updateBase()
	fmt.Println(triangle)

	// Using Pointers Directly without functions
	pointerToTriangle := &triangle
	// This is one way
	(*pointerToTriangle).base = 1
	fmt.Println(triangle)
	// This is another way. But both does the same thing.
	// This is the recommended way.
	pointerToTriangle.base = 2
	fmt.Println(triangle)
}

func PointersToStructUsingVar() {
	triangle := Triangle1{10, 4}
	fmt.Println(triangle)

	// Using Pointers with Struct Functions
	triangle.updateBase()
	fmt.Println(triangle)
	// Using Pointers Directly without functions
	// pointerToTriangle := &triangle
	var pointerToTriangle *Triangle1 = &triangle
	// This is one way
	(*pointerToTriangle).base = 1
	fmt.Println(triangle)
	// This is another way. But both does the same thing.
	// This is the recommended way.
	pointerToTriangle.base = 2
	fmt.Println(triangle)
}
