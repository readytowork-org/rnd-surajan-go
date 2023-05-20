package variables_constants

import "fmt"

func SimpleMultipleDeclaration() {
	var magicNum, powerLevel int32
	magicNum, powerLevel = 2048, 9001

	fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)
	// Prints: "magicNum is: 2048 powerLevel is: 9001"

	// We can also do this
	var marks, total float32 = 55, 80
	fmt.Println("Marks:", marks, "Out of:", total)
}

func MultipleDeclarationUsingShort() {
	// Here, "amount" has value 10 and "unit" has value "doll hairs"
	amount, unit := 10, "doll hairs"

	fmt.Println(amount, unit, ", that's expensive...")
	// Prints: "10 doll hairs , that's expensive..."
}
