package main

import (
	"fmt"
)

/*configuration
1-increase channel
2-decrease channel
3-enter channel number
4-exit */

func channel() {
	//channel := []string{"SunTv", "Ktv", "VijayTv", "ZeeTv"}
	var channel1 = make(map[int]string)
	channel1[1] = "SunTV"
	channel1[2] = "KTV"
	channel1[3] = "VijayTV"
	channel1[4] = "ZeeTV"
	var t, input int
	current := 1
	check := 0
	fmt.Println("channel", current, channel1[current])
	fmt.Scanf("%d", &t)
	for t != 4 {
		if t == 1 && check == 0 {
			current += 1
			fmt.Println(current, channel1[current])
			if current > 4 {
				current = 1
				check = 1
				fmt.Println(current, channel1[current])
				continue
			}
		} else if t == 2 && check == 0 {
			current--
			fmt.Println(current, channel1[current])
			if current == 0 {
				current = 4
				check = 1
				fmt.Println(current, channel1[current])
				continue
			}
		} else if t == 3 {
			fmt.Scanf("%d", &input)
			fmt.Println(channel1[input])
		} else if t == 4 {
			break
		}

		fmt.Scanf("%d", &t)
		check = 0

	}

}
func description() {
	fmt.Println("configuration \n1-increase channel \n2-decrease channel \n3-enter channel number \n4-exit ")
}
func main() {
	//var employee = make(map[string]int)
	//var ptr [3]*int //creating array with pointer type
	description()
	channel()
}
