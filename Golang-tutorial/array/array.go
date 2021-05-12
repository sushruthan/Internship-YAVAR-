package main

import (
	"fmt"
	"strings"
)

const MAX int = 3

func main() {
	greetings := []string{"Hello", "world!"} //creating a slice
	fmt.Println(strings.Join(greetings, " "))
	arrayint()
	x := make3D(2, 2, 3) //calling function

	x[1][0][2] = 9 //assigning value of index 1,0,2 for 3d array
	fmt.Println(x)
	fmt.Printf("%T\n", x) //printing the type of variable

	/* an int array with 5 elements */
	var balance = []int{1000, 2, 3, 17, 50}
	var avgr float32 = getAverage(balance, 5)

	/* pass array as an argument */
	//avgr = getAverage(balance, 5)
	/* local variable definition */
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	/* calling a function to swap the values.
	* &a indicates pointer to a ie. address of variable a and
	* &b indicates pointer to b ie. address of variable b.
	 */
	swap(&a, &b) //send pointer value

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)

	/* output the returned value */
	fmt.Printf("Average value is: %f\n ", avgr)
	basicPointer()
	pointarray()
	ptrptr()
}
func arrayint() {
	var n [10]int //n is an array of 10 integers
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100 //setting array value to hundred
	}

	/* output each array element's value */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

func make3D(m, n, p int) [][][]float64 { //passing value to the function and returning array
	buf := make([]float64, m*n*p) //creating 1d array
	fmt.Println(buf)
	x := make([][][]float64, m) //creating 3d array
	fmt.Println(x)
	for i := range x {
		x[i] = make([][]float64, n) //creating 2d array
		for j := range x[i] {       //assigning values for the 2d array
			x[i][j] = buf[:p:p]
			buf = buf[p:]
		}
	}
	return x
}

func getAverage(arr []int, size int) float32 { //passing array in a function
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum / size)
	return avg
}
func basicPointer() {
	var a int = 20   /* actual variable declaration */
	var ip *int = &a /* pointer variable declaration */

	//ip = &a /* store address of a in pointer variable*/

	fmt.Printf("Address of a variable: %x\n", &a)

	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)
	var ptr *int //declaring nill pointer

	fmt.Printf("The value of ptr is : %x %T\n", ptr, ptr)
	if ptr != nil {
		fmt.Println("not a nill pointer")
	} /* succeeds if p is not nil */
	if ptr == nil {
		fmt.Println("ptr is a nill pointer")
	} /* succeeds if p is null */
}
func pointarray() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int //creating array with pointer type

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* assign the address of integer. */
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *ptr[i])
	}
}
func ptrptr() { //function to assign pointer value to a pointer
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* take the address of var */
	ptr = &a

	/* take the address of ptr using address of operator & */
	pptr = &ptr

	/* take the value using pptr */
	fmt.Printf("Value of a = %d\n", a)
	fmt.Printf("Value available at *ptr = %d,%d\n", *ptr, ptr)
	fmt.Printf("Value available at **pptr = %d,%d\n", **pptr, pptr)
}

func swap(x *int, y *int) { //pass pointer as parameter
	var temp int
	temp = *x /* save the value at address x */
	*x = *y   /* put y into x */
	*y = temp /* put temp into y */
}
