package arrays

import "fmt"

func AppendToSlice() {
	myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
	fmt.Println(myTutors)
	// Append a new Tutor named "Josh" into myTutors.
	myTutors = append(myTutors, "Josh")
	fmt.Println(myTutors)
}
