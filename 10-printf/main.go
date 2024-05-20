package main

import (
	"fmt"
	os "os"
)

func main() {
	age := 20
	fmt.Printf("I am %d years old.\n", age)
	firstName, lastName := "John", "Doe"
	fmt.Printf("My name is %s %s.\n", firstName, lastName)
	claim := false
	fmt.Printf("This statement is %t.\n", claim)
	double := 3.14
	fmt.Printf("The value of pi is to 1dp is %.1f.\n", double)
	fmt.Printf("The type of double is %T.\n", double)
	fmt.Printf("\"hi\"\n")

	if len(os.Args) != 3 {
		fmt.Println("Please provide a first and last name.")
		return
	}

	fn, ln := os.Args[1], os.Args[2]
	msg := "Input names: %s %s\n"
	fmt.Printf(msg, fn, ln)
}