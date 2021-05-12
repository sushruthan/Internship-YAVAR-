package main

import (
	"fmt"
	"math"
)

func fibonaci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

// recursive function for
// printing all numbers
// upto the number n
func print_one(n int) {

	// if the number is positive
	// print the number and call
	// the second function
	if n >= 0 {
		fmt.Println("In first function:", n)
		// call to the second function
		// which calls this first
		// function indirectly
		print_two(n - 1)
	}
}

func print_two(n int) {

	// if the number is positive
	// print the number and call
	// the second function
	if n >= 0 {
		fmt.Println("In second function:", n)
		// call to the first function
		print_one(n - 1)
	}
}
func conversionExplict() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("Value of mean : %f\n", mean)
}

//define an interface
type Shape interface {
	area() float64
}

//define a circle
type Circle struct {
	x, y, radius float64
}

// define a rectangle
type Rectangle struct {
	width, height float64
}

// define a method for circle (implementation of Shape.area())
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// define a method for rectangle (implementation of Shape.area())
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

// define a method for shape
func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", fibonaci(i))
	}

	// passing a positive
	// parameter which prints all
	// numbers from 1 - 10
	print_one(10)

	// this will not print
	// anything as it does not
	// follow the base case
	print_one(-1)
	conversionExplict()
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}
	fmt.Printf("Circle area: %T\n", circle)
	fmt.Printf("Rectangle area: %f\n", rectangle)
	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
