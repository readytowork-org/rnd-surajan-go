package structs

import "fmt"

type Cake struct {
	typeOfCake string
	weight     int // Weight in grams
}

func ArrayOfStructs() {
	// Array of Structs.
	cakes := []Cake{{"Chocolate", 1000}, {"Carrot", 500}, {"Ice Cream", 2000}}
	// Displaying weight of "chocolate" cake.
	fmt.Println(cakes[0].weight)
	// Updating weight of "carrot" cake.
	cakes[1].weight = 1500
	fmt.Println(cakes)
}
