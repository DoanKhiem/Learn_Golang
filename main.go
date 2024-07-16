package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("tim")
	sayGreeting("jenny")
	sayBye("tim")

	cycleNames([]string{"tim", "jenny"}, sayGreeting)
	cycleNames([]string{"tim", "jenny"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("Circle 1 is %.3f and circle 2 is %.3f \n", a1, a2)
}

//func main() {
//	age := 45
//
//	fmt.Println(age <= 50)
//	fmt.Println(age >= 50)
//	fmt.Println(age == 45)
//	fmt.Println(age != 50)
//
//	if age < 30 {
//		fmt.Println("Age is less than 30")
//	} else if age < 40 {
//		fmt.Println("Age is less than 40")
//	} else {
//		fmt.Println("Age is 50 or more")
//	}
//
//	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}
//
//	for index, value := range names {
//		if index == 1 {
//			fmt.Println("continuing at pos", index)
//			continue
//		}
//		if index > 2 {
//			fmt.Println("breaking at pos", index)
//			break
//		}
//		fmt.Printf("The value at pos %v is %v \n", index, value)
//	}
//}

//func main() {
//	x := 0
//	for x < 5 {
//		fmt.Println("Value of x is:", x)
//		x++
//	}
//
//	for i := 0; i < 5; i++ {
//		fmt.Println("Value of i is:", i)
//	}
//
//	names := []string{"mario", "luigi", "yoshi", "peach"}
//
//	for i := 0; i < len(names); i++ {
//		fmt.Println(names[i])
//	}
//
//	for index, value := range names {
//		fmt.Printf("The index is %v and the value is %v \n", index, value)
//	}
//
//	for _, value := range names {
//		fmt.Printf("The value is %v \n", value)
//		value = "new string"
//	}
//
//	fmt.Println(names)
//}

//func main() {
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

//	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
//
//	sort.Ints(ages)
//	fmt.Println(ages)
//
//	index := sort.SearchInts(ages, 30)
//	fmt.Println(index)
//
//	names := []string{"yoshi", "mario", "peach", "bowser"}
//
//	sort.Strings(names)
//	fmt.Println(names)
//
//	fmt.Println(sort.SearchStrings(names, "bowser"))
//}
