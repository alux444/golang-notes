package main

import (
	"fmt"
	"os"
	"strconv"
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

	if len(os.Args) != 4 {
		fmt.Println("Expected 3 arguments, got", len(os.Args)-1)
		return
	}

	fn, ln := os.Args[1], os.Args[2]
	msg := "Input names: %s %s\n"
	fmt.Printf(msg, fn, ln)

	//convert third argument to float
	arg, error := strconv.ParseFloat(os.Args[3], 64)
	fmt.Printf("The third argument is %f\n", arg)
	fmt.Println("The error code is: ", error)
}