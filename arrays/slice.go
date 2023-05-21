package arrays

import "fmt"

func CreateSlice() {
	// create array below
	lottery := []int{5, 48, 32, 1, 6}
	fmt.Println(lottery)
}

func CreateSliceFromArray() {
	myTutors := [4]string{"Kirsty", "Mishell", "Jose", "Neil"}
	fmt.Println(myTutors)
	// Creating a slice using "myTutors" array.
	tutors := myTutors[:]
	// Creating a new slice
	subjects := []string{"Go", "Java", "Python"}
	// Displaying both the slices
	fmt.Println(tutors)
	fmt.Println(subjects)
	// We can access slices the same way we access arrays.
	fmt.Printf("Ms. %v teaches %v", tutors[0], subjects[1])
}

func CreateSliceFromArrayByTargetingElements() {
	array := [5]int{2, 5, 7, 1, 3}
	// This is a slice of the whole array
	sliceVersion := array[:]
	fmt.Println(sliceVersion)
	// [2 5 7 1 3]
	// This is a slice containing the elements at indices 2, 3, and 4
	partialSlice := array[2:5]
	fmt.Println(partialSlice)
	// [7 1 3]
}
