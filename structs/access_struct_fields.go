package structs

import "fmt"

type Restaurant struct {
	name             string
	typeOfRestaurant string
	yearEstablished  int
}

func AccessStructFields() {
	// Add your code here.
	restaurant := Restaurant{"Codecademy Steakhouse", "Japanese", 2011}
	fmt.Println(restaurant)
	// Accessing fields.
	fmt.Println(restaurant.name)
	fmt.Println(restaurant.typeOfRestaurant)
	fmt.Println(restaurant.yearEstablished)
	// Updating fields.
	restaurant.name = "Skillsoft Steakhouse"
	restaurant.yearEstablished = 2022
	fmt.Println(restaurant)
}
