package main

import (
	"fmt"
	"os"
)

func main() {
	jsonFile, err := os.Open("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	fmt.Println(jsonFile)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
}
