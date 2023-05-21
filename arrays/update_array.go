package arrays

import "fmt"

func UpdateArrayElements() {
	// I have 3 dogs, Frida, Fido, and Jeff
	myDogs := [3]string{"Frida", "Fedo", "Jegf"}
	fmt.Println(myDogs)
	myDogs[1] = "Fido"
	myDogs[2] = "Jeff"
	fmt.Println(myDogs)
}
