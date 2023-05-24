package type_checking

import "fmt"

func UsingTVerb() {
	gpa := 4.2
	yesNo := true
	var petrolPrice float32 = 180.25
	// Type checking
	fmt.Printf("Type of variable 'gpa' is: %T", gpa)
	fmt.Printf("\nType of variable 'yesNo' is: %T", yesNo)
	fmt.Printf("\nType of variable 'petrolPrice' is: %T", petrolPrice)
}
