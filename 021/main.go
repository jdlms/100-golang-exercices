// Exercise: MAP
// Create a map of ints to strings
// 1 should resolve A
// 2 should resolve B
// 3 should resolve C

package main

import "fmt"

func main() {

	aMap := make(map[int]string)

	aMap[1] = "A"
	aMap[2] = "B"
	aMap[3] = "C"

	fmt.Println(aMap)

}
