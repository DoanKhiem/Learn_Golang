package main

import (
	"fmt"
	"math"
)

// shape interface
type shape interface {
	area() float64
	circumf() float64
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

// square method
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return 4 * s.length
}

// circle method
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Println("Area: ", s.area())
	fmt.Println("Circumference: ", s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 15},
		circle{radius: 5},
		circle{radius: 12},
		square{length: 25},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("----")
	}
}
