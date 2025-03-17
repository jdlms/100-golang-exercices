// Exercise: User input
// Using only the fmt package, ask a user for it's name and then for it's surname
// Store it in 2 variables called "name" and "surname"
// After user has entered the data, print it out

// Tip: https://pkg.go.dev/fmt#hdr-Scanning

package main

import "fmt"

func main() {
	firstName := ""
	lastName := ""
	fmt.Println("What is your first name?")
	fmt.Scanln(&firstName)

	fmt.Println("What is your last name?")
	fmt.Scanln(&lastName)

	fmt.Printf("Hello, %s %s\n", firstName, lastName)
}
