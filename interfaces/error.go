package interfaces

import "fmt"

// Interface collecting all the methods.
type Shape2 interface {
	area() float32
	// "perimeter()" method which has not been implemented by "Rectangle" struct anywhere. So, this should throw an error.
	perimeter() float32
}

// "Rectangle2" struct.
type Rectangle2 struct {
	length, breadth float32
}

// "area2()" function for "Rectangle" struct.
func (r Rectangle) area2() float32 {
	return r.length * r.breadth
}

// Access method of the interface to access "area2()"
func calculate2(s Shape, shapeName string) {
	fmt.Printf("Area of %v: %.1f\n", shapeName, s.area())
}

// ❗ Throws Error: When a struct does not implement all the methods of the given interface
func ErrorInInterface() {
	rect := Rectangle2{10, 20}
	// Will throw error.
	// ❗ Might even show a red line below "rect" in Editor.
	calculate2(rect, "Rectangle")
}
