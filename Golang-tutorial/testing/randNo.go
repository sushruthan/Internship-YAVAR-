package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Now().UnixNano())
	// for {
	// 	fmt.Println(funran())
	// }
	fmt.Printf("%T", funran())

}
func funran() int {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	ran := y1.Intn(time.Now().Hour())
	//return y1.Intn(time.Now().UnixNano())
	return ran
}
