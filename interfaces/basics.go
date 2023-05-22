package interfaces

import "fmt"

// Interface collecting all the methods.
type Shape interface {
	// "area()" method returns a "float32" value.
	area() float32
	// "dimensions()" method does not return a value.
	dimensions()
}

// "Rectangle" struct.
type Rectangle struct {
	length, breadth float32
}

// "area()" function for "Rectangle" struct.
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

// "dimensions()" function for "Rectangle" struct.
func (r Rectangle) dimensions() {
	fmt.Printf("Length: %f and Breadth: %f.", r.length, r.breadth)
}

// Access method of the interface to access "area()"
func calculate(s Shape) {
	fmt.Println("Area:", s.area())
}

// Access method of the interface to access "dimensions()"
func getDimensions(s Shape) {
	s.dimensions()
}

func ShowInterface() {
	rect := Rectangle{10, 20}
	// Calling functions with struct variable "rect()".
	calculate(rect)
	getDimensions(rect)
}
