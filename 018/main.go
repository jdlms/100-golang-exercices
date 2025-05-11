// Exercise: Write a list of 5 countries to a file
// Tip: use the "os" package

package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("countries")
	if err != nil {
		fmt.Printf("Error %s", err)
	}

	defer file.Close()

	file.WriteString("Peru\nFrance\nEcuador\nGermany\nCanada")
}
