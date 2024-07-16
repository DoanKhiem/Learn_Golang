package main

import (
	"fmt"
	"sort"
)

func main() {
	//// Strings
	//var nameOne string = "mario"
	//var nameTwo = "luigi"
	//var nameThree string
	//
	//fmt.Println(nameOne, nameTwo, nameThree)
	//
	//nameOne = "peach"
	//nameThree = "bowser"
	//
	//fmt.Println(nameOne, nameTwo, nameThree)
	//
	//// khai bao ngan gon
	//nameFour := "yoshi"
	//
	//fmt.Println(nameFour)
	//
	//// Numbers
	//var ageOne int = 20
	//var ageTwo = 30
	//
	//ageThree := 40
	//
	//fmt.Println(ageOne, ageTwo, ageThree)

	// Bits & Memory
	//var numOne int8 = 25
	//var numTwo int8 = -128
	//var numThree uint8 = 255

	//var scoreOne float32 = 25.98
	//var scoreTwo float64 = 25.9832233232323323232332323232323232323
	//scoreThree := 42.5

	// Print
	//age := 35
	//name := "shaun"
	//
	//fmt.Print("Hello, ")
	//fmt.Print("world!\n")
	//fmt.Print("New line \n")
	//
	//// Println
	//fmt.Println("Hello, World!")
	//fmt.Println("Goodbye, World!")
	//fmt.Println("My age is", age, "and my name is", name)
	//
	//// Printf (formatted strings)
	//fmt.Printf("My age is %v and my name is %v \n", age, name)
	//fmt.Printf("My age is %q and my name is %q \n", age, name)
	//fmt.Printf("age is of type %T \n", age)
	//fmt.Printf("You scored %0.1f \n", 225.55)
	//
	//// Sprintf (save formatted strings)
	//var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	//fmt.Println("The saved string is:", str)

	// Arrays
	//var ages [3]int = [3]int{20, 25, 30}
	//var ages = [3]int{20, 25, 30}
	//
	//names := [4]string{"yoshi", "mario", "peach", "bowser"}
	//
	//fmt.Println(ages, len(ages))
	//fmt.Println(names, len(names))
	//
	//// Slices
	//var scores = []int{100, 50, 60}
	//scores[1] = 25
	//scores = append(scores, 85)
	//
	//fmt.Println(scores, len(scores))
	//
	//// Slice ranges
	//rangeOne := names[1:3]
	//rangeTwo := names[2:]
	//rangeThree := names[:3]
	//
	//fmt.Println(rangeOne, rangeTwo, rangeThree)
	//
	//rangeOne = append(rangeOne, "koopa")
	//fmt.Println(rangeOne)

	//greeting := "hello there friends!"
	//fmt.Println(strings.Contains(greeting, "hello"))
	//fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	//fmt.Println(strings.ToUpper(greeting))
	//fmt.Println(strings.Index(greeting, "ll"))
	//fmt.Println(strings.Split(greeting, " "))
	//
	//// The original string remains unchanged
	//fmt.Println("Original string value: ", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser"))
}
