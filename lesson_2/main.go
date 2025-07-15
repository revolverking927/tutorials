package main

import "fmt"

func main() {
	//String variables
	var nameOne string = "Alice"
	var nameTwo = "Bob"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Charlie"
	nameThree = "David"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Eve"
	fmt.Println(nameFour)

	// Integer variables
	var ageOne int = 30
	var ageTwo = 25
	ageThree := 20
	fmt.Println(ageOne, ageTwo, ageThree)

	// Bits and memory
	var ageFour int8 = 22
	var ageFive int16 = 28912
	var ageSix uint16 = 256
	fmt.Println(ageFour, ageFive, ageSix)

	var scoreOne float32 = 95.5
	var scoreTwo float64 = 88.75
	fmt.Println(scoreOne, scoreTwo)
}
