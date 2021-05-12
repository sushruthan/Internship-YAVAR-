package main

import (
	"fmt"
	"time"
)

func main() {
	heartBeat()
	time.Sleep(time.Second * 50)
}

func heartBeat() {
	for range time.Tick(time.Second * 5) {
		fmt.Println("Foo")
	}
}
