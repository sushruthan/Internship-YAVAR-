package main

import (
	"fmt"
	"strconv"
)

var (
	j    int    = 5         // var uppercase letter can be exported
	Name string = "haggrid" //lowercase can be used inside package scope
)

func main() {
	var i, b int //block delaration
	i = 63
	b = 2
	j = 3 //shadowing
	var k float32
	var g string

	fmt.Printf("%T\n", j)
	k = float32(i) //explicit conversion
	g = string(i)
	fmt.Println("hello", i, b, j, g, Name)
	fmt.Printf("%v %T %T\n", j, k, g)

	g = strconv.Itoa(i)
	fmt.Printf("check %T %v\n", g, g)
	sample(5.5)

}
func sample(i float64) {

	fmt.Println("sushruthan", i)
}
