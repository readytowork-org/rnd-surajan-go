package arrays

import "fmt"

func LengthAndCapacity() {
	slice := []string{"Fido", "Fifi", "FruFru"}
	// The slice begins at length 3 and capacity 3
	fmt.Println(slice, len(slice), cap(slice))
	// [Fido Fifi FruFru] 3 3
	slice = append(slice, "FroFro")
	// After appending an element when the slice is at capacity
	// The slice will double in capacity, but increase its length by 1
	fmt.Println(slice, len(slice), cap(slice))
	// [Fido Fifi FruFru FroFro] 4 6
}
