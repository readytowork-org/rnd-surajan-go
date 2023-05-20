package randomizing

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomizingProblem() {
	// This will print a random number between 0 and 99.
	// However, on some devices it might only going to print one value (in Codecademy's Editor, it printed only 81) everytime as we have not set out seed value.
	// Here, the default seed value Go uses is 1.
	fmt.Println(rand.Intn(100))

	// But, on our local computer it seems to show random numbers. Hence, this might not be a problem on local computer.
}

// But, in some machines, it might be a problem. The solution is below.
func RandomizingSolution() {
	// Here, we are generating a seed number using the current time generated in milliseconds.
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
}
