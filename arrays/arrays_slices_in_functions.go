package arrays

import "fmt"

// Changes to the array will only be local to the function
func changeArrayFirstElement(array [4]int, value int) {
	array[0] = value
}

func ChangeArray() {
	lottery := [4]int{5, 48, 32, 1}
	changeArrayFirstElement(lottery, 100)

	fmt.Println(lottery)
}

// Changes to the slice parameter will be permanent
func changeSliceFirstElement(slice []int, value int) {
	if len(slice) > 0 {
		slice[0] = value
	}
}

func ChangeSlice() {
	lottery := []int{5, 48, 32, 1}
	changeSliceFirstElement(lottery, 100)

	fmt.Println(lottery)
}

func changeLastElement(slice []string, value string) {
	slice[len(slice)-1] = value
	fmt.Println(slice)
}

func ChangeAnotherSlice() {
	myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
	changeLastElement(myTutors, "Bobby")
}
