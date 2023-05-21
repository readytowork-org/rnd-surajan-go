package loops

import "fmt"

func BreakContinue() {
	for count := 0; count < 20; count++ {
		// IF COUNT EQUALS 8 BELOW THIS LINE, THIS ITERATION WILL BE SKIPPED
		if count == 8 {
			continue
		}
		// IF COUNT EQUALS 15 BELOW THIS LINE, THIS ENTIRE LOOP WILL END
		if count == 15 {
			break
		}
		fmt.Println(count)
	}
}
