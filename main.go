package main

import "fmt"

func main() {
	// Strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	// khai bao ngan gon
	nameFour := "yoshi"

	fmt.Println(nameFour)

	// Numbers
	var ageOne int = 20
	var ageTwo = 30

	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// Bits & Memory
	//var numOne int8 = 25
	//var numTwo int8 = -128
	//var numThree uint8 = 255

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 25.9832233232323323232332323232323232323
	scoreThree := 42.5
}
