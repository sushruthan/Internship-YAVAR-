package main

import "fmt"

func main() {
	fmt.Printf("Hello\tWorld!")
	const s = `First line
	Second line`
	fmt.Println(s)
	fmt.Println("`" + "foo" + "`") // Prints: `foo`
	const g = "\tFirst line\n" +
		"Second line"
	fmt.Println(g)
	fmt.Println("\"food\"") // Prints: "foo" double quote
	fmt.Println("\v canned food")
	fmt.Println("\\caf\u00e9") // Prints string: \café
	fmt.Printf("%c", '\u00e9') // Prints character: é
	fmt.Println("\b money")
}
