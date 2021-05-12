// // Generating Random Numbers in Golang
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func randomno() int {

// 	// Seeding - Go provides a method,
// 	// Seed(see int64), that allows you
// 	// to initialize this default sequence.

// 	// Implementation is slow
// 	// to make it faster
// 	// rand.Seed(time.Now().UnixNano())
// 	// is added. Seed is the current time,
// 	// converted to int64 by UnixNano.
// 	// Gives constantly changing numbers
// 	x1 := rand.NewSource(time.Now().UnixNano())
// 	y1 := rand.New(x1)
// 	then := time.Now()

// 	//return y1.Intn(time.Now().UnixNano())
// 	Println(y1.Intn(then.Hour()))

// }

// // func randomNumber() {
// // 	go heartBeat()
// // 	time.Sleep(time.hour * 10)
// // }

// // func heartBeat() {
// //     now := time.Now()
// // 	for range time.Tick(time.Minute * 5) {

// // 		tags := "Random Number"
// // 		fields := randomno()

// // 		acc.AddFields("Random number", fields, tags, now)
// // 	}
// // }

// // func init() {
// // 	inputs.Add("randomNumber", func() telegraf.Input {
// // 		return &randomNumber{}
// // 	})
// // }
// func main(){
// 	fmt.Println(time.Now().UnixNano())
// }
