// Exercise: Random
// Create a program that mimics a dice roll (a 6-face dice)
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var num = random(1, 6)
	fmt.Printf("%d\n", num)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
