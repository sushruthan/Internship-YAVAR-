package main

import (
	"fmt"
)

//Boolean type
//Numeric type
//String type
//Derived type
//
var u uint8 = 255 //0-255
var u1 uint16     //0-655355
var u2 uint32     //0-4294967295
var u3 uint64     //0-18446744073709551615

var i int8 = 'A'
var i1 int16 = 'A'
var i2 int32
var i3 int64

//machine dependent
var mac1 uint
var mac2 int = 45
var mac3 uintptr

//float
var f float32
var f1 float64

//complex
var c complex64
var c1 complex128 = complex(1, 5)

//boolean
var b bool

//string
var s string = "hello"

//byte is uint8
var by byte = 'A' //takes ascii values

//rune is a int32 data type
var r rune = 'A'

/* declaring multiple variable
a, b, c := 1, 2, 3
var a, b, c string
const a, b, c = 3, 4, "foo"  // a = 3, b = 4, c = "foo", untyped integer and string constants
const u, v float32 = 0, 3    // u = 0.0, v = 3.0
x, y := "Hello", 10   // x is an instance of `string`, y is of type `int`
*/
/*var (
	a = 12
	b = 3
	enableA = false

	food = "chocolate"
	myvar float64
	anothervar float64 = 2.4
 )

//Also works for const

 const (
   xconst    = 5
   boolconst = false
 )
*/
//function to print data types
func constOne() {
	const (
		hello string  = "Hello"
		hi            //automatically assigns previous type and value
		date  float32 = 3.4
	)
	const (
		a = iota // a == 0
		b = iota // b == 1
		c = iota // c == 2
		d        // d == 3 (implicitely d = iota)
	)
	/*
		const a = iota // 0
		const b = iota // 0
		const (
		  c = iota // 0
		  d = iota // 1
		)
		const (
		  e = iota // 0
		  f = iota // 1
		)
	*/
	//var location *string = &hello //constants dont have address
	const day = "monday" //untyped constant
	fmt.Println("Printing constants \n", hello, date, hi, day)
	fmt.Println("Printing constants \n", a, b, c, d)

}

/*const (
	a = iota + 1 // a == 1
	_            // (implicitly _ = iota + 1 )
	b            // b == 3 (implicitly b = iota + 1 )
	c            // c == 4 (implicitly c = iota + 1 )
)*/
func printTypes() {
	fmt.Printf("%T, %T, %T ,%T, %v, %v, %v, %v \n", u, u1, u2, u3, u, u1, u2, u3)
	fmt.Printf("%T, %T, %T ,%T, %v, %v, %v, %v  \n", i, i1, i2, i3, i, i1+i1, i2, i3)
	fmt.Printf("%T, %T, %v, %v \n", f, f1, f, f1)
	fmt.Printf("%T, %T, %v,%v \n", c, c1, c, c1)
	fmt.Printf("%T, %v \n", b, b)
	fmt.Printf("%T, %v \n", s, s)
	fmt.Printf("%T, %T, %T, %v, %v, %v \n", mac1, mac2, mac3, mac1, mac2, mac3)
}
func conversion() {
	var number int = 65
	var name string = "Sam"
	var decimal float32 = 3.6
	fmt.Printf("%T, %T, %T, %v, %v, %v \n", number, name, decimal, number, name, decimal)
	con1 := float32(number) //int to float conersion
	con2 := string(number)  //int to string conversion
	con3 := int(decimal)    //float32 to int
	fmt.Printf("converted values: %T, %T, %T,%v, %v, %v \n", con1, con2, con3, con1, con2, con3)
}
func main() {
	b = true
	c = complex(2, 9)   //assigning values to complex data type variable
	var p1 *int16 = &i1 //"&" used to assign the pointer value
	var p2 *int = &mac2

	//mac3 = p2 //uintptr doesnt take ptr values
	printTypes()                                               //calling function
	fmt.Printf("Byte %T, %v+200 =, %v \n", by, by, (by + 200)) //value exceeds
	fmt.Printf("Rune %T, %v, \n", r, (r + 65))
	fmt.Printf("%T,%T, %v, %v \n", p1, p2, *p1, p2) //by using "*" can retrive the value using pointer
	conversion()                                    //calling function
	//fmt.Println(mac3)
	constOne()
}
