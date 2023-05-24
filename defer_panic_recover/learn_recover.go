package defer_panic_recover

import "fmt"

func handlePanic() {
	// "recover()" will detect if panic occurs or not. And if it does, it catches the panic message and assigns to "recoverMsg".
	recoverMsg := recover()
	if recoverMsg != nil {
		fmt.Println("Recover:", recoverMsg)
	}
}

func division(num1, num2 int) {
	// Execute "handlePanic()" even after panic occurs.
	defer handlePanic()

	if num2 == 0 {
		panic("Cannot divide by zero")
	} else {
		fmt.Printf("Result: %d\n", num1/num2)
	}
}

func LearnRecover() {
	division(4, 2)
	division(4, 0)
	division(4, 3)
}
