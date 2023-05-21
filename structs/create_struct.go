package structs

import "fmt"

// Country struct goes here
type Country struct {
	name      string
	capital   string
	latitude  float32
	longitude float32
}

func CreateStruct() {
	var france Country
	fmt.Println(france)
}
