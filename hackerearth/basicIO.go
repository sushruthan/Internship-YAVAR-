package main

import (
	"fmt"
)

func main() {
	var T int
	// upper := 0
	// lower := 0
	// ugly := 0
	fmt.Scanf("%d", &T)
	a := make([]string, T)
	for i := 0; i < T; i++ {
		fmt.Scanf("%s", &a[i])
	}
	for i := 0; i < T; i++ {
		for _, s := range a[i] {

			//fmt.Println(s, string(s))
			d := string(s)
			r := []rune(d)
			fmt.Println(r)
		}
	}
	// for i := 0; i < T; i++ {
	// 	upper = 0
	// 	lower = 0
	// 	ugly = 0
	// 	for _, s := range a[i] {
	// 		fmt.Printf("%c\n", s)
	// 		ascii := int(s)
	// 		s := string(r)
	// 		fmt.Println(s)
	// 		if (unicode.IsUpper(s) == true) && (unicode.ToUpper(s) == 'A') && (unicode.ToUpper(s) == 'E') && (unicode.ToUpper(s) == 'I') && (unicode.ToUpper(s) == 'O') && (unicode.ToUpper(s) == 'U') {
	// 			upper++
	// 			fmt.Println("Upper", s)
	// 		} else if (unicode.IsLower(s) == true) && (unicode.ToUpper(s) == 'a') && (unicode.ToUpper(s) == 'e') && (unicode.ToUpper(s) == 'i') && (unicode.ToUpper(s) == 'o') && (unicode.ToUpper(s) == 'u') {
	// 			lower++
	// 			fmt.Println("lower", s)
	// 		} else if (unicode.IsUpper(s) == true) && (unicode.ToUpper(s) != 'A') && (unicode.ToUpper(s) != 'E') && (unicode.ToUpper(s) != 'I') && (unicode.ToUpper(s) != 'O') && (unicode.ToUpper(s) != 'U') {
	// 			ugly++//[72][65][69][73][79][85]
	// 		}

	// 	}
	// 	if lower == 0 && upper > 0 {
	// 		fmt.Println("lovely string")
	// 	} else if lower > 0 && upper == 0 {
	// 		fmt.Println("lovely string")
	// 	} else if ugly > 0 && lower > 0 || upper > 0 {
	// 		fmt.Println("lovely string")
	// 	} else {
	// 		fmt.Println("ugly string")
	// 		fmt.Println(ugly, lower, upper)
	// 	}
	// }
}
