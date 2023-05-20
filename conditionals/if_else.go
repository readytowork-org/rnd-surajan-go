package conditionals

import "fmt"

func UsingIf() {
	alarmRinging := true
	// This works.
	if alarmRinging {
		fmt.Println("Turn off the alarm!!")
	}

	// This also works.
	if alarmRinging {
		fmt.Println("Turn off the alarm!!")
	}
}

func UsingElse() {
	heistReady := false

	if heistReady {
		fmt.Println("Let's go!")
	} else {
		fmt.Println("Act normal.")
	}
}

func IfElseWithComparisonOp() {
	lockCombo := "2-35-19"
	robAttempt := "2-35-19"

	if lockCombo == robAttempt {
		fmt.Println("The vault is now opened.")
	}

	vaultAmt := 2356468

	if vaultAmt >= 200000 {
		fmt.Println("We're going to need more bags.")
	}
}

func IfElseWithLogicalOp() {
	rightTime := true
	rightPlace := true

	// With "&&" operator
	if rightTime && rightPlace {
		fmt.Println("We're outta here!")
	} else {
		fmt.Println("Be patient...")
	}

	enoughRobbers := false
	enoughBags := true

	// With "||" operator
	if enoughRobbers || enoughBags {
		fmt.Println("Grab everything!")
	} else {
		fmt.Println("Grab whatever you can!")
	}
}
