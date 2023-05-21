package structs

import "fmt"

type Point struct {
	x int
	y int
}

// "Circle" struct uses a "point" field of type of "Point" struct.
type Circle struct {
	point  Point
	radius int
}

func NormalNestedStructs() {
	circle := Circle{Point{4, 5}, 2}

	// We can access like this.
	fmt.Println(circle.point.x)
	fmt.Println(circle)

}

// "Circle1" struct has an anonymous field of type of "Point" struct.
// Here, we didn't use a variable like: "point".
type Circle1 struct {
	Point
	radius int
}

func AnonymousNestedStruct() {
	circle := Circle1{Point{4, 5}, 2}

	// To access, we don't have to call "circle.point.x", we can directly use "circle.x"
	fmt.Println(circle.x)
	fmt.Println(circle)

}
