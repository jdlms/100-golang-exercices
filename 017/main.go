// Exercise: Read a file
// Tip: use the "io/ioutil" package

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./read.go")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(data))
}
