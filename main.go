package main

import (
	"fmt"
	// 💡 Note: Import the packages from a module(folder) which you wanna play with and run "go run ." command.
	// 👇 Below are some of the examples on how to do so.
	// 🙂 One way to import.
	"rnd-surajan-go/greet"
	// 😎 Another way to import. Here the name "variables_constants" is too long, so we gave an alias: "vc".
	vc "rnd-surajan-go/variables_constants"
)

func main() {
	fmt.Println("Welcome to Golang.")
	// 💡 Note: The exported function from any of the given packages should have a CamelCase convention. Like: SayHello, not sayHello.
	greet.SayHello()
	greet.SayGM()
	vc.Variables()
}
