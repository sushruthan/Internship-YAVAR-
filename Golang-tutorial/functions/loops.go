package main

import (
	"fmt"
	"math"
	"time"
)

func ifopp() {
	var check int = 100

	//check the boolean condition
	if (check == 10) || (check >= 100) { //using
		// if condition is true then print
		if check < 20 { //nested if loops
			// if condition is true then print
			fmt.Printf("check is less than 20\n")
		} else {
			// if condition is false then print
			fmt.Printf("check is not less than 20\n")
		}
	} else if check == 20 {
		// if else if condition is true
		fmt.Printf("Value of check is 20\n")
	} else if check == 30 {

		fmt.Printf("Value of check is 30\n")
	} else {

		fmt.Printf("None of the values is matching\n")
	}
	fmt.Printf("Exact value of check is: %d\n", check)
}
func switchtype() { //switch case to determine the type of the varaible
	var x interface{} //Empty interfaces are used by code that handles values of unknown type.

	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x :%T", i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is float64")
	case func(int) float64:
		fmt.Printf("x is func(int)")
	case bool, string:
		fmt.Printf("x is bool or string")
	default:
		fmt.Printf("don't know the type")
	}
}
func switchTime() {
	today := time.Now()
	fmt.Println("time", today)
	switch today.Day() {
	case 5:
		fmt.Println("Clean your house.")
		fallthrough //to force the execution flow to fall through the successive case block
	case 10:
		fmt.Println("Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 25:
		fmt.Println("Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func forloop() {
	var b int = 15
	var a int
	numbers := [6]int{1, 2, 3, 5}

	/* for loop execution */
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}
	for a < b {
		a++
		fmt.Printf("value of a: %d\n", a)
	}
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}
}
func primenumber() {
	/* local variable definition */
	var i, j int

	for i = 2; i < 20; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // if factor found, not prime,terminates the loop
			}
		}
		if j > (i / j) {
			fmt.Printf("%d is prime\n", i) //print prime number
		}
	}
}
func continue1() {
	/* local variable definition */
	var a int = 10

	/* do loop execution */
	for a < 20 {
		if a == 15 {
			/* skip the iteration */
			a = a + 1
			continue
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
}
func goto1() {
	/* local variable definition */
	var a int = 10

	/* do loop execution */
LOOP:
	for a < 20 {
		if a == 15 {
			/* skip the iteration */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
}
func swap(x, y string) (string, string) { //function with parameter to return value
	return y, x
}
func getSequence() func() int { //declaring anonymous function
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func init() {
	fmt.Println("Welcome to init() function") //used to declare global value
}
func main() {

	ifopp()       //calling function
	switchtype()  //calling function
	forloop()     //calling function
	primenumber() //calling function
	continue1()   //calling function
	goto1()       //calling function
	a, b := swap("Mahesh", "Kumar")
	switchTime()
	fmt.Println(a, b)

	/* declare a function variable */
	getSquareRoot := func(x float64) float64 { //calling function as variable
		return math.Sqrt(x)
	}

	fmt.Println("square root of 9 is ", getSquareRoot(9)) //calling function
	/* nextNumber is now a function with i as 0 */
	nextNumber := getSequence()

	/* invoke nextNumber to increase i by 1 and return the same */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	var i interface{} //Empty interfaces are used by code that handles values of unknown type.
	describe(i)
	i = 42
	describe(i)
	i = "hello"
	describe(i)

}
