/*
First, you mash in a random large number to start with. Then, repeatedly do the following:

If the number is divisible by 3, divide it by 3.

If it's not, either add 1 or subtract 1 (to make it divisible by 3), then divide it by 3.

The game stops when you reach "1".

*/

package main

import (
	"fmt"
	"strconv"
	"math/rand"
)

func main() {
	trecimate(rand.Int())
}

func trecimate(n int) int {
	num := n
	for {
		fmt.Printf("Have %s, ", strconv.Itoa(num))
		if (num % 3) == 0 {
			fmt.Println("dividing by 3 \n")
			num = num / 3
		} else {
			fmt.Println("subtracting 1 \n")
			num = num -1
		}
		if( num == 1) {
			fmt.Println("Got one!")
			break
		}
	}

	return num
}
