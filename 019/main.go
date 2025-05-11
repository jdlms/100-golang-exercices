// Exercise: Rename a file from name1 to name2
// Tip: use the "os" package

package main

import (
	"fmt"
	"os"
)

func main() {
	var src, dest string
	src = "name1.txt"

	_, err := os.Create(src)
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	dest = "renamed-file.txt"

	err = os.Rename(src, dest)
	if err != nil {
		fmt.Printf("Error renaming: %s", err)
	}
}
