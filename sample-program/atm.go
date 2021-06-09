package main

import (
	"fmt"

)

func main() {
	var T int
	fmt.Scanf("%d", &T)
	a := make([]string, T)
	for i := 0; i < T; i++ {
		fmt.Scanf("%s", &a[i])
	}
