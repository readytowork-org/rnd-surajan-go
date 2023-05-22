package interfaces

import "fmt"

// Interface collecting all the methods.
type Shape1 interface {
	area() float32
}

// "Rectangle1" struct.
type Rectangle1 struct {
	length, breadth float32
}

// "area1()" function for "Rectangle1" struct.
func (r Rectangle) area1() float32 {
	return r.length * r.breadth
}

// "Triangle" struct.
type Triangle struct {
	base, height float32
}

// "area()" function for "Triangle" struct.
func (t Triangle) area() float32 {
	return (t.base * t.height) / 2
}

// Access method of the interface to access "area()"
func calculate1(s Shape1, shapeName string) {
	fmt.Printf("Area of %v: %.1f\n", shapeName, s.area())
}

func InterfaceWithMultipleStructs() {
	rect := Rectangle{10, 20}
	tri := Triangle{5, 10}
	// Each instance can call the calculate function by itself.
	calculate1(rect, "Rectangle")
	calculate1(tri, "Triangle")
}
