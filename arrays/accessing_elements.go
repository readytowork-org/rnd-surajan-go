package arrays

import "fmt"

func AccessArrayElements() {
	triangleAngles := [3]int{30, 60, 90}
	fmt.Println(triangleAngles[2])
	// Sum
	sum := triangleAngles[0] + triangleAngles[1] + triangleAngles[2]
	fmt.Println(sum)
}
