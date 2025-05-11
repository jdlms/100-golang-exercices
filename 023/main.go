// Exercise: Random
// Generate a random number from te range [-50, +50]
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var num = random(-50, 50)
	fmt.Printf("%d\n", num)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
